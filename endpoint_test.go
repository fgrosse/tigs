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

	It("should not generate anything if endpoint is abstract", func() {
		buf := &bytes.Buffer{}
		output.Writer = buf
		ep := endpoint{
			Name:   "GetStuff",
			Method: "GET", URL: "/stuff",
			Abstract: true,
			Parameters: []parameter{
				{Name: "s", TypeString: "string"},
				{Name: "i", TypeString: "int"},
			},
		}

		ep.generate(output, "TestClient")
		Expect(buf.Bytes()).To(BeEmpty())
	})

	It("should generate GET operations", func() {
		ep := endpoint{
			Name:   "GetStuff",
			Method: "GET", URL: "/stuff",
			Parameters: []parameter{
				{Name: "s", TypeString: "string"},
				{Name: "i", TypeString: "int"},
			},
		}

		ep.generate(output, "TestClient")
		Expect(output.Writer).To(ContainCode(`
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
			Name:   "CreateStuff",
			Method: "POST", URL: "/stuff",
			Parameters: []parameter{
				{Name: "s", TypeString: "string", Location: "query"},
				{Name: "b", TypeString: "bool", Location: "json"},
				{Name: "i", TypeString: "int", Location: "json"},
			},
		}

		ep.generate(output, "TestClient")
		Expect(output.Writer).To(ContainCode(`
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
