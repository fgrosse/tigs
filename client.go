package main

import (
	"io"
	"strings"
)

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
}

func (c client) generateType(out *formattableWriter) {
	c.generateTypeComment(out)

	out.printf(`type %s struct {`, c.Name)
	out.printf(`	BaseURL *url.URL`)
	out.printf(`	Client  tigshttp.Client`)
	out.printf(`}`)
}

func (c client) generateTypeComment(out *formattableWriter) {
	if c.Description == "" {
		return
	}

	comment := c.Description
	if len(comment) < len(c.Name) || comment[:len(c.Name)] != c.Name {
		comment = c.Name + ": " + comment
	}

	comment = strings.TrimSpace(comment)
	comment = strings.Replace(comment, "\n", "\n// ", -1)
	out.printf(`// %s`, comment)
}

func (c client) generateFactoryFunction(w io.Writer) {
	out := &formattableWriter{w}

	out.printf(``)
	out.printf(`func New%s(baseURL string) (*%s, error) {`, c.Name, c.Name)
	out.printf(`	u, err := url.Parse(baseURL)`)
	out.printf(`	if err != nil {`)
	out.printf(`		return nil, fmt.Errorf("invalid base URL for new %s: %%s", err)`, c.Name)
	out.printf(`	}`)
	out.printf(``)
	out.printf(`	return &%s{`, c.Name)
	out.printf(`		BaseURL: u,`)
	out.printf(`		Client: http.DefaultClient,`)
	out.printf(`	}, nil`)
	out.printf(`}`)
}

func (c client) imports() []string {
	imports := []string{
		"fmt",
		"net/http",
		"net/url",
		"github.com/fgrosse/tigs/tigshttp",
	}

	if c.containsJSONEndpoints() {
		imports = append(imports, "encoding/json", "bytes", "io/ioutil")
	}

	return imports
}

func (c client) containsJSONEndpoints() bool {
	for _, ep := range c.Endpoints {
		if ep.hasJSONParameters() {
			return true
		}
	}

	return false
}
