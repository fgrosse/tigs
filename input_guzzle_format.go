package main

import (
	"fmt"
	"os"
	"sort"

	"encoding/json"
	"gopkg.in/yaml.v2"
)

const (
	GuzzleYAML = "guzzle-yaml"
	GuzzleJSON = "guzzle-json"
)

func init() {
	registeredUnmarshallers[GuzzleYAML] = &guzzleServiceDescriptionUnmarshaller{GuzzleYAML}
	registeredUnmarshallers[GuzzleJSON] = &guzzleServiceDescriptionUnmarshaller{GuzzleJSON}
}

type guzzleServiceDescriptionUnmarshaller struct {
	typ string
}

type guzzleServiceDescription struct {
	Name, Description, Version string

	Operations map[string]guzzleEndpointDescription
	Imports    []string
}

type guzzleEndpointDescription struct {
	Summary, Method, URI string

	Abstract   bool
	Parameters map[string]parameter
}

func (u *guzzleServiceDescriptionUnmarshaller) Unmarshal(input []byte, c *client) (err error) {
	description := guzzleServiceDescription{}

	switch u.typ {
	case GuzzleYAML:
		input = sanitizeYAML(input)
		err = yaml.Unmarshal(input, &description)
	case GuzzleJSON:
		err = json.Unmarshal(input, &description)
	default:
		return fmt.Errorf("unknown guzzleServiceDescriptionUnmarshaller type %q", u.typ)
	}

	if err != nil {
		return err
	}

	for _, i := range description.Imports {
		// TODO i should actually be interpreted as lying relative to the input file
		importedFile, err := os.Open(i)
		if err != nil {
			return fmt.Errorf("could not open imported file: %s", err)
		}

		importedDef := new(client)
		err = newDecoder(u.typ, importedFile).decode(importedDef)
		if err != nil {
			return fmt.Errorf("could not decode imported file: %s", err)
		}

		// this should result in another guzzleServiceDescription
		// now apply the other guzzleServiceDescription (omit empty values)
		if importedDef.Name != "" {
			c.Name = importedDef.Name
		}
		if importedDef.APIVersion != "" {
			c.APIVersion = importedDef.APIVersion
		}
		if importedDef.Package != "" {
			c.Package = importedDef.Package
		}
		if importedDef.Description != "" {
			c.Description = importedDef.Description
		}


		c.Endpoints = append(c.Endpoints, importedDef.Endpoints...)
		logDebug("Imported %d new endpoints from file %q (total %d)", len(importedDef.Endpoints), i, len(c.Endpoints))
	}

	description.translateInto(c)
	return nil
}

func (d guzzleServiceDescription) translateInto(c *client) {
	if d.Name != "" {
		c.Name = d.Name
	}

	c.Description = d.Description
	c.APIVersion = d.Version
	c.Endpoints = append(c.Endpoints, d.translateOperations()...)
}

func (d guzzleServiceDescription) translateOperations() []endpoint {
	endpoints := make([]endpoint, len(d.Operations))
	epNames := make([]string, len(d.Operations))

	// make endpoint order deterministic to simplify tests
	i := 0
	for name := range d.Operations {
		epNames[i] = name
		i++
	}
	sort.Strings(epNames)

	i = 0
	for _, epName := range epNames {
		o := d.Operations[epName]
		endpoints[i] = d.translateOperation(o)
		endpoints[i].Name = epName
		i++
	}

	return endpoints
}

func (d guzzleServiceDescription) translateOperation(op guzzleEndpointDescription) *endpoint {
	if op.Abstract {
		return nil
	}

	ep := &endpoint{
		Description: op.Summary,
		Method:      op.Method,
		URL:         op.URI,
		Parameters:  make([]parameter, len(op.Parameters)),
	}

	i := 0
	for _, name := range op.orderedParameterNames() {
		p := op.Parameters[name]
		ep.Parameters[i] = parameter{
			Name:        name,
			Description: p.Description,
			Type:        p.Type,
			Location:    p.Location,
			Required:    p.Required,
		}
		i++
	}

	return ep
}

func (op guzzleEndpointDescription) orderedParameterNames() []string {
	keys := make([]string, len(op.Parameters))
	i := 0
	for name := range op.Parameters {
		keys[i] = name
		i++
	}
	sort.Strings(keys)
	return keys
}
