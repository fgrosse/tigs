package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
	"sort"
)

type decoder interface {
	Decode(*client) error
}

type baseDecoder struct {
	io.Reader
	logger io.Writer
}

type yamlDecoder struct {
	baseDecoder
}

func newYAMLDecoder(input io.Reader) *yamlDecoder {
	return &yamlDecoder{
		baseDecoder: baseDecoder{
			Reader: input,
			logger: ioutil.Discard,
		},
	}
}

func (d *yamlDecoder) Decode(c *client) error {
	input, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}

	type yamlParameter struct {
		Name, Description, Type, Location string
		Required                          bool
	}

	type yamlEndpoint struct {
		Name, Description, Method, URL string

		Abstract bool
		Summary  string // alternative for "description"
		URI      string // alternative for "url"

		Parameters map[string]yamlParameter
	}

	var unmarshalled struct {
		Name, Description string
		APIVersion        string `yaml:"apiVersion"`
		Version           string // alternative for "apiVersion"

		Endpoints  map[string]yamlEndpoint
		Operations map[string]yamlEndpoint // alternative for endpoints
	}

	input = d.sanitizeInput(input)
	err = yaml.Unmarshal(input, &unmarshalled)
	if err != nil {
		return err
	}

	unpackEndpoints := func(eps map[string]yamlEndpoint) []endpoint {
		r := make([]endpoint, len(eps))
		epNames := make([]string, len(eps))
		i := 0
		for name := range eps {
			epNames[i] = name
			i++
		}
		sort.Strings(epNames)

		i = 0
		for _, epName := range epNames {
			ep := eps[epName]
			r[i] = endpoint{
				Name:        epName,
				Abstract:    ep.Abstract,
				Description: firstOf(ep.Description, ep.Summary),
				Method:      ep.Method,
				URL:         firstOf(ep.URL, ep.URI),
				Parameters:  make([]parameter, len(ep.Parameters)),
			}

			parameterNames := make([]string, len(ep.Parameters))
			j := 0
			for name := range ep.Parameters {
				parameterNames[j] = name
				j++
			}
			sort.Strings(parameterNames)

			j = 0
			for _, name := range parameterNames {
				p := ep.Parameters[name]
				r[i].Parameters[j] = parameter{
					Name:        name,
					Description: p.Description,
					TypeString:  p.Type,
					Location:    p.Location,
					Required:    p.Required,
				}
				j++
			}

			i++
		}

		return r
	}

	c.Name = unmarshalled.Name
	c.Description = unmarshalled.Description
	c.APIVersion = firstOf(unmarshalled.APIVersion, unmarshalled.Version)

	if len(unmarshalled.Endpoints) > 0 {
		c.Endpoints = unpackEndpoints(unmarshalled.Endpoints)
	} else {
		c.Endpoints = unpackEndpoints(unmarshalled.Operations)
	}

	return nil
}

func (d *yamlDecoder) sanitizeInput(input []byte) []byte {
	var sanitizedInput = &bytes.Buffer{}

	line := &bytes.Buffer{}
	lineBeginning := true
	for _, c := range input {
		switch c {
		case '\n':
			if strings.TrimSpace(line.String()) != "" {
				sanitizedInput.Write(append(line.Bytes(), '\n'))
				line.Reset()
				lineBeginning = true
			}
		case '\t':
			if lineBeginning {
				line.WriteString("    ")
			} else {
				line.WriteByte(c)
			}
		case ' ':
			line.WriteByte(c)
		default:
			lineBeginning = false
			line.WriteByte(c)
		}
	}

	sanitizedInput.Write(line.Bytes())

	s := sanitizedInput.Bytes()
	fmt.Fprintf(d.logger, "Sanitized input is:\n%s", string(s))
	return s
}

func firstOf(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}

	return ""
}
