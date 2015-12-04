package main

import "fmt"

// A client holds all information necessary to generate a go client for an HTTP web service.
type client struct {
	// Name is the name of the generated go type for this client.
	Name string

	// Description is a textual summary of this client which is used when generating documentation.
	Description string

	// APIVersion is the version of the API this client expects to communicate with.
	APIVersion string

	// Package is the name of the package that the generated client code will be defined in.
	Package string

	// Endpoints is a list of HTTP endpoints that are available over this client.
	Endpoints []endpoint

	// imports is the go code that is generated for the import ( ... ) block
	Imports string
}

func (c client) containsJSONEndpoints() bool {
	for _, ep := range c.Endpoints {
		if ep.hasParameterWithType("json") {
			return true
		}
	}

	return false
}

func (c client) containsPostfieldEndpoints() bool {
	for _, ep := range c.Endpoints {
		if ep.hasParameterWithType("postField") {
			return true
		}
	}

	return false
}

func (c client) Validate() error {
	if c.Package == "" {
		return fmt.Errorf("missing package")
	}

	if len(c.Endpoints) == 0 {
		return fmt.Errorf("no endpoints")
	}

	for _, ep := range c.Endpoints {
		err := ep.Validate()
		if err != nil {
			return fmt.Errorf("invalid endpoint %q: %s", ep.Name, err)
		}
	}

	return nil
}
