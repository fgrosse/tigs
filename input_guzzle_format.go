package main

import (
	"sort"
	"fmt"

	"gopkg.in/yaml.v2"
	"encoding/json"
)

func init() {
	registeredUnmarshallers["guzzle-yaml"] = &guzzleServiceDescriptionUnmarshaller{"yaml"}
	registeredUnmarshallers["guzzle-json"] = &guzzleServiceDescriptionUnmarshaller{"json"}
}

type guzzleServiceDescriptionUnmarshaller struct {
	typ string
}

func (u *guzzleServiceDescriptionUnmarshaller) Unmarshal(input []byte, c *client) (err error) {
	description := guzzleServiceDescription{}

	switch u.typ {
	case "yaml":
		input = sanitizeYAML(input)
		err = yaml.Unmarshal(input, &description)
	case "json":
		err = json.Unmarshal(input, &description)
	default:
		return fmt.Errorf("unknown guzzleServiceDescriptionUnmarshaller type %q", u.typ)
	}

	if err != nil {
		return err
	}

	description.translateInto(c)
	return nil
}

type guzzleServiceDescription struct {
	Name, Description, Version string

	Operations map[string]guzzleEndpointDescription
}

type guzzleEndpointDescription struct {
	Summary, Method, URI string

	Abstract   bool
	Parameters map[string]parameter
}

func (d guzzleServiceDescription) translateInto(c *client) {
	c.Name = d.Name
	c.Description = d.Description
	c.APIVersion = d.Version
	c.Endpoints = d.translateOperations()
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

func (d guzzleServiceDescription) translateOperation(op guzzleEndpointDescription) endpoint {
	ep := endpoint{
		Abstract:    op.Abstract,
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
