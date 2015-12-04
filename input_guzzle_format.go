package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

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
	Summary, URI string
	HTTPMethod   string `json:"httpMethod" yaml:"httpMethod"`

	Abstract   bool
	Extends    string
	Parameters map[string]parameter
}

func (u *guzzleServiceDescriptionUnmarshaller) Unmarshal(input []byte, c *client) (err error) {
	description := new(guzzleServiceDescription)

	switch u.typ {
	case GuzzleYAML:
		input = sanitizeYAML(input)
		err = yaml.Unmarshal(input, description)
	case GuzzleJSON:
		err = json.Unmarshal(input, description)
	default:
		return fmt.Errorf("unknown guzzleServiceDescriptionUnmarshaller type %q", u.typ)
	}

	if err != nil {
		return err
	}

	if err = description.translateInto(c); err != nil {
		return err
	}

	for _, i := range description.Imports {
		Debug.Printf("Importing referenced file %q\n", i)
		// TODO i should actually be interpreted as lying relative to the input file
		importedFile, err := os.Open(i)
		if err != nil {
			return fmt.Errorf("could not open imported file: %s", err)
		}

		importedDef := new(client)
		err = newDecoder(u.typ, importedFile).decode(importedDef, settings{inheritance: false}) // TODO pass decoder options and tell it not to resolve inheritance stuff on imports immediately)
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
		Debug.Printf("Imported %d new endpoints from file %q (total %d)\n", len(importedDef.Endpoints), i, len(c.Endpoints))
	}

	return nil
}

func (d *guzzleServiceDescription) translateInto(c *client) error {
	if d.Name != "" {
		c.Name = d.Name
	}

	c.Description = d.Description
	c.APIVersion = d.Version
	c.Endpoints = append(c.Endpoints, d.translateOperations()...)

	return nil
}

func (d *guzzleServiceDescription) translateOperations() []endpoint {
	endpoints := []endpoint{}

	if len(d.Operations) == 0 {
		Debug.Printf("Found no endpoints to decode")
		return endpoints
	}

	// make endpoint order deterministic to simplify tests
	epNames := []string{}
	for name := range d.Operations {
		epNames = append(epNames, name)
	}
	sort.Strings(epNames)

	for _, epName := range epNames {
		ep := d.translateOperation(d.Operations[epName])
		ep.Name = epName
		endpoints = append(endpoints, ep)
	}

	Debug.Printf("Successfully decoded %d endpoint(s): %q", len(epNames), epNames)
	return endpoints
}

func (d *guzzleServiceDescription) translateOperation(op guzzleEndpointDescription) endpoint {
	ep := endpoint{
		Description: op.Summary,
		Method:      op.HTTPMethod,
		Extends:     op.Extends,
		Abstract:    op.Abstract,
		URI:         op.URI,
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
