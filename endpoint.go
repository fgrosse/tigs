package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// An endpoint represents an operation of a service that can be accessed by clients via an URL.
type endpoint struct {
	ClientName  string
	Name        string
	Abstract    bool
	Extends     string
	Description string
	Method      string
	URL         string
	Parameters  []parameter
}

type formattableWriter struct {
	io.Writer
}

func (w *formattableWriter) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}

func (ep endpoint) Generate() string {
	if ep.Method == "" {
		ep.Method = "GET"
	}

	args := []string{}
	for _, p := range ep.Parameters {
		args = append(args, fmt.Sprintf(p.Name)+" "+p.generatedType())
	}

	buf := &bytes.Buffer{}
	out := &formattableWriter{buf}
	out.printf(``)
	out.printf(`func (c *%s) %s(%s) (*http.Response, error) {`, ep.ClientName, ep.Name, strings.Join(args, ", "))
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

	return buf.String()
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
