package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"fmt"
)

var _ = Describe("endpoint", func() {
	var output *formattableWriter
	BeforeEach(func() {
		output = &formattableWriter{&bytes.Buffer{}}
		fmt.Fprintln(output, "package tigs_test") // generate a package name so the generated code will have no syntax errors
	})

	It("should generate GET operations", func() {
		ep := endpoint{
			Name:       "GetStuff",
			ClientName: "TestClient",
			Method:     "GET", URL: "/stuff",
			Parameters: []parameter{
				{Name: "s", Type: "string"},
				{Name: "i", Type: "int"},
			},
		}

		Expect("package tigs_test" + ep.Generate()).To(ContainCode(`
			func (c *TestClient) GetStuff(s string, i int) (*http.Response, error) {
				u, err := c.BaseURL.Parse("/stuff")
				if err != nil {
					return nil, err
				}

				u.Query().Add("s", s)
				u.Query().Add("i", fmt.Sprintf("%d", i))

				req := tigshttp.NewRequest("GET", u)
				return c.Client.Do(req)
			}
		`))
	})

	It("should generate POST operations", func() {
		ep := endpoint{
			Name:       "CreateStuff",
			ClientName: "TestClient",
			Method:     "POST", URL: "/stuff",
			Parameters: []parameter{
				{Name: "s", Type: "string", Location: "query"},
				{Name: "b", Type: "bool", Location: "json"},
				{Name: "i", Type: "int", Location: "json"},
			},
		}

		Expect("package tigs_test" + ep.Generate()).To(ContainCode(`
			func (c *TestClient) CreateStuff(s string, b bool, i int) (*http.Response, error) {
				u, err := c.BaseURL.Parse("/stuff")
				if err != nil {
					return nil, err
				}

				u.Query().Add("s", s)

				data, err := json.Marshal(map[string]interface{}{
					"b": b,
					"i": i,
				})

				if err != nil {
					return nil, fmt.Errorf("could not marshal body for CreateStuff: %s", err)
				}

				req := tigshttp.NewRequest("POST", u)
				req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
				req.ContentLength = len(data)
				req.Header.Set("Content-Type", "application/json")

				return c.Client.Do(req)
			}
		`))
	})
})
