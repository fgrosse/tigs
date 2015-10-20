package tigs

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"fmt"
	"io"
)

var _ = Describe("Endpoint", func() {
	var output io.Writer
	BeforeEach(func() {
		output = &bytes.Buffer{}
		fmt.Fprintln(output, "package tigs_test") // generate a package name so the generated code will have no syntax errors
	})
	It("should generate GET operations", func() {
		ep := Endpoint{
			Name:   "GetStuff",
			Method: "GET", URL: "/stuff",
			Parameters: []Parameter{
				{Name: "s", Type: "string"},
				{Name: "i", Type: "int"},
			},
		}

		Expect(ep.Generate(output, "TestClient")).To(Succeed())
		Expect(output).To(ContainCode(`
			func (c *TestClient) GetStuff(s string, i int) (*http.Response, error) {
				u, err := c.BaseURL.Parse("/stuff")
				if err != nil {
					return nil, err
				}

				u.Query().Add("s", s)
				u.Query().Add("i", fmt.Sprintf("%d", i))

				return c.Client.Get(u.String())
			}
		`))
	})

	It("should generate POST operations", func() {
		ep := Endpoint{
			Name:   "CreateStuff",
			Method: "POST", URL: "/stuff",
			Parameters: []Parameter{
				{Name: "s", Type: "string", Location: "query"},
				{Name: "b", Type: "bool", Location: "json"},
				{Name: "i", Type: "int", Location: "json"},
			},
		}

		Expect(ep.Generate(output, "TestClient")).To(Succeed())
		Expect(output).To(ContainCode(`
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

				return c.Client.Post(u.String(), "application/json", bytes.NewBuffer(data))
			}
		`))
	})
})
