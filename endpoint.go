package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type endpoint struct {
	ClientName  string
	Name        string
	Abstract    bool
	Extends     string
	Description string
	Method      string
	URI         string
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

	ep.generateURLCode(out)

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

	if ep.hasParameterWithType("json") {
		out.printf("	data, err := json.Marshal(map[string]interface{}{")
		for _, p := range ep.Parameters {
			if p.Location == "json" {
				out.printf("\t\t%q: %s,", p.Name, p.Name)
			}
		}
		out.printf("	})")
		out.printf("")
		out.printf("	if err != nil {")
		out.printf("		return nil, fmt.Errorf(\"could not marshal body for %s: %%s\", err)", ep.Name)
		out.printf("	}")
		out.printf("")
	}

	if ep.hasParameterWithType("postField") {
		out.printf("	data := url.Values{")
		for _, p := range ep.Parameters {
			if p.Location == "postField" {
				out.printf("\t\t%q: {%s},", p.Name, p.stringCode())
			}
		}
		out.printf("	}")
		out.printf("")
	}

	out.printf("\treq := tigshttp.NewRequest(%q, u)", ep.Method)

	if ep.hasParameterWithType("json") {
		out.printf("\treq.Body = ioutil.NopCloser(bytes.NewBuffer(data))")
		out.printf("\treq.ContentLength = len(data)")
		out.printf("\treq.Header.Set(\"Content-Type\", \"application/json\")")
		out.printf("")
	}

	if ep.hasParameterWithType("postField") {
		out.printf("\treq.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))")
		out.printf("\treq.ContentLength = len(data)")
		out.printf("\treq.Header.Set(\"Content-Type\", \"application/x-www-form-urlencoded\")")
		out.printf("")
	}

	out.printf("	return c.Client.Do(req)")
	out.printf("}")

	return buf.String()
}

func (ep endpoint) generateURLCode(out *formattableWriter) {
	if ep.hasParameterWithType("uri") == false {
		out.printf(`	u, err := c.BaseURL.Parse(%q)`, ep.URI)
		return
	}

	out.printf(`	u, err := tigshttp.ExpandURITemplate(%q, map[string]interface{}{`, ep.URI)
	for _, p := range ep.Parameters {
		if p.Location == "uri" {
			out.printf(`		%q: %s,`, p.Name, p.Name)
		}
	}
	out.printf(`	})`)
	fmt.Fprintln(out)
}

func (ep endpoint) hasQueryParameters() bool {
	for _, p := range ep.Parameters {
		if p.Location == "" || p.Location == "query" {
			return true
		}
	}

	return false
}

func (ep endpoint) hasParameterWithType(t string) bool {
	for _, p := range ep.Parameters {
		if p.Location == t {
			return true
		}
	}

	return false
}

func (ep endpoint) Validate() error {
	if ep.Name == "" {
		return fmt.Errorf("missing name")
	}

	if ep.URI == "" {
		return fmt.Errorf("missing URI")
	}

	l := ""
	for _, p := range ep.Parameters {
		err := p.Validate()
		if err != nil {
			return fmt.Errorf("invalid parameter %q: %s", p.Name, err)
		}

		if l != "" && p.Location != l {
			return fmt.Errorf("incompatible parameter locations")
		}

		l = p.Location
	}

	return nil
}
