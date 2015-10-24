package main

import (
	"fmt"
	"strings"
)

// An endpoint represents an operation of a service that can be accessed by clients via an URL.
type endpoint struct {
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
	Parameters []parameter
}

func (ep endpoint) generate(out *formattableWriter, clientName string) {
	args := []string{}
	for _, p := range ep.Parameters {
		args = append(args, fmt.Sprintf(p.Name)+" "+p.generatedType())
	}

	out.printf(``)
	out.printf(`func (c *%s) %s(%s) (*http.Response, error) {`, clientName, ep.Name, strings.Join(args, ", "))
	out.printf(`	u, err := c.BaseURL.Parse(%q)`, ep.URL) // TODO check what happens if baseURL = foobar/v1/ and ep path = /blup
	out.printf(`	if err != nil {`)
	out.printf(`		return nil, err`)
	out.printf(`	}`)

	if ep.hasQueryParameters() {
		out.printf(``)
	}

	for _, p := range ep.Parameters {
		if p.Location == "" || p.Location == "query" {
			out.printf("	u.Query().Add(%q, %s)", p.Name, p.stringCode())
		}
	}

	if len(ep.Parameters) > 0 {
		out.printf(``)
	}

	if ep.hasJSONParameters() {
		out.printf("\tdata, err := json.Marshal(map[string]interface{}{")
		for _, p := range ep.Parameters {
			if p.Location == "json" {
				out.printf("\t\t\"%s\": %s,", p.Name, p.Name)
			}
		}
		out.printf("	})")
		out.printf("")
		out.printf("	if err != nil {")
		out.printf("		return nil, fmt.Errorf(\"could not marshal body for %s: %%s\", err)", ep.Name)
		out.printf("	}")
		out.printf("")
	}

	out.printf("\treq := tigshttp.NewRequest(%q, u)", ep.Method)
	if ep.hasJSONParameters() {
		out.printf("\treq.Body = ioutil.NopCloser(bytes.NewBuffer(data))")
		out.printf("\treq.ContentLength = len(data)")
		out.printf("\treq.Header.Set(\"Content-Type\", \"application/json\")")
		out.printf("")
	}

	out.printf("	return c.Client.Do(req)")
	out.printf("}")
}

func (ep endpoint) hasQueryParameters() bool {
	for _, p := range ep.Parameters {
		if p.Location == "" || p.Location == "query" {
			return true
		}
	}

	return false
}

func (ep endpoint) hasJSONParameters() bool {
	for _, p := range ep.Parameters {
		if p.Location == "json" {
			return true
		}
	}

	return false
}
