package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
	"unicode/utf8"
)

type decoder struct {
	unmarshaller
	io.Reader

	err error // err is the deferred error that might have happened in the call to newDecoder()
}

type settings struct {
	inheritance bool
	forceExport bool
}

var registeredUnmarshallers = map[string]unmarshaller{}

type unmarshaller interface {
	Unmarshal(input []byte, c *client) (err error)
}

func newDecoder(inputType string, input io.Reader) decoder {
	u, isDefined := registeredUnmarshallers[inputType]
	if !isDefined {
		return decoder{err: fmt.Errorf("unknown input type %q", inputType)}
	}

	return decoder{u, input, nil}
}

func (d decoder) decode(c *client, s settings) error {
	if d.err != nil {
		return d.err
	}

	input, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}

	err = d.Unmarshal(input, c)
	if err != nil {
		return err
	}

	if s.forceExport {
		for i := range c.Endpoints {
			c.Endpoints[i].Name = upperCaseFirst(c.Endpoints[i].Name)
			c.Endpoints[i].Extends = upperCaseFirst(c.Endpoints[i].Extends)
		}
	}

	if s.inheritance {
		return newEndpointTree(c.Endpoints).process()
	}

	return nil
}

func upperCaseFirst(s string) string {
	if s == "" {
		return ""
	}

	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func sanitizeYAML(input []byte) []byte {
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
	return s
}
