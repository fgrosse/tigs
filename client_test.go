package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"fmt"
)

var _ = Describe("client", func() {
	var output *formattableWriter
	BeforeEach(func() {
		output = &formattableWriter{&bytes.Buffer{}}
		fmt.Fprintln(output, "package tigs_test\n") // generate a package name so the generated code will have no syntax errors
	})

	Describe("generating code", func() {
		It("should define the client and use the given package name", func() {
			c := client{Name: "MyClient"}

			c.generateType(output)
			Expect(output.Writer).To(ContainCode(`
				type MyClient struct {
					BaseURL *url.URL
					Client  tigshttp.Client
				}
			`))
		})

		It("should generate a type comment", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "MyClient is awesome"}

			c.generateType(output)
			Expect(output.Writer).To(ContainCode(`
				// MyClient is awesome
				type MyClient struct {
			`))
		})

		It("should always start type comments with the type name", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "This is some description"}

			c.generateType(output)
			Expect(output.Writer).To(ContainCode(`
				// MyClient: This is some description
				type MyClient struct {
			`))
		})

		It("should prepend each new line with `//` to mark it as comment", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "This is some description\nover multiple lines"}

			c.generateType(output)
			Expect(output.Writer).To(ContainCode(`
				// MyClient: This is some description
				// over multiple lines
				type MyClient struct {
			`))
		})

		It("should provide a New* function", func() {
			c := client{Name: "TestClient"}

			c.generateFactoryFunction(output)
			Expect(output.Writer).To(ContainCode(`
				func NewTestClient(baseURL string) (*TestClient, error) {
					u, err := url.Parse(baseURL)
					if err != nil {
						return nil, fmt.Errorf("invalid base URL for new TestClient: %s", err)
					}

					return &TestClient{
						BaseURL: u,
						Client: http.DefaultClient,
					}, nil
				}
			`))
		})

		Describe("retrieving a list of imports", func() {
			It("should return all packages necessary for the factory function", func() {
				c := client{}

				Expect(c.imports()).To(ContainElement("net/url"))
				Expect(c.imports()).To(ContainElement("net/http"))
				Expect(c.imports()).To(ContainElement("fmt"))
			})

			It("should return all packages necessary if there are json parameters", func() {
				c := client{
					Endpoints: []endpoint{
						{Method: "POST", Name: "Do", Parameters: []parameter{{Name: "p", Location: "json"}}},
					},
				}

				Expect(c.imports()).To(ContainElement("bytes"))
				Expect(c.imports()).To(ContainElement("encoding/json"))
				Expect(c.imports()).To(ContainElement("io/ioutil"))
				Expect(c.imports()).To(ContainElement("github.com/fgrosse/tigs/tigshttp"))
			})
		})
	})
})
