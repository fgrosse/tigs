package tigs

import (
	"io"
	"fmt"
	"strings"
)

// An Endpoint represents an operation of a service that can be accessed by clients via an URL.
type Endpoint struct {
	// Name is the symbolic name of this endpoint.
	// It is used in code generate and does not represent any part of the URL that is actually used.
	// Every endpoint must always have a Name or it is considered invalid.
	Name string

	// Description is an optional textual representation of this endpoint.
	// It is used to generate documentation and is not meant to be used automatically.
	Description string

	// Method is the HTTP method of this endpoint.
	Method string

	// URL is the URL under which the endpoint can be reached.
	URL string

	// Parameters is the list of parameters that are used to create the HTTP request.
	Parameters []Parameter
}

func (ep Endpoint) Generate(w io.Writer, clientName string) error {
	out := &formattableWriter{w}

	args := []string{}
	for _, p := range ep.Parameters {
		args = append(args, fmt.Sprintf(p.Name)+" "+p.GeneratedType())
	}

	out.printf(``)
	out.printf(`func (c *%s) %s(%s) (*http.Response, error) {`, clientName, ep.Name, strings.Join(args, ", "))
	out.printf(`    u, err := c.BaseURL.Parse(%q)`, ep.URL) // TODO check what happens if baseURL = foobar/v1/ and ep path = /blup
	out.printf(`    if err != nil {`)
	out.printf(`        return nil, err`)
	out.printf(`    }`)

	if ep.HasQueryParameter() {
		out.printf(``)
	}

	for _, p := range ep.Parameters {
		if p.Location == "" || p.Location == "query" {
			out.printf("\tu.Query().Add(%q, %s)", p.Name, p.StringCode())
		}
	}

	if len(ep.Parameters) > 0 {
		out.printf(``)
	}

	switch ep.Method {
	case "GET":
		out.printf("\treturn c.Client.Get(u.String())")
	case "POST":
		out.printf("\tdata, err := json.Marshal(map[string]interface{}{")
		for _, p := range ep.Parameters {
			if p.Location == "json" {
				out.printf("\t\t\"%s\": %s,", p.Name, p.Name) // TODO order parameters and format indent
			}
		}
		out.printf("\t})")
		out.printf("")
		out.printf("\tif err != nil {")
		out.printf("\t\treturn nil, fmt.Errorf(\"could not marshal body for %s: %%s\", err)", ep.Name)
		out.printf("\t}")
		out.printf("")
		out.printf("\treturn c.Client.Post(u.String(), \"application/json\", bytes.NewBuffer(data))")
	default:
		panic("NOT IMPLEMENTED")
	}

	out.printf(`}`)
	return nil
}

func (ep Endpoint) HasQueryParameter() bool {
	for _, p := range ep.Parameters {
		if p.Location == "" || p.Location == "query" {
			return true
		}
	}

	return false
}

func (ep Endpoint) HasJSONParameter() bool {
	for _, p := range ep.Parameters {
		if p.Location == "json" {
			return true
		}
	}

	return false
}
