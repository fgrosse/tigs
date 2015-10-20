package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"fmt"
	"io"
)

var _ = Describe("Service Client", func() {
	var (
		output io.Writer
		client ServiceClient
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
		fmt.Fprintln(output, "package tigs_test") // generate a package name so the generated code will have no syntax errors
	})

	Describe("generating code", func() {
		It("should define the client and use the given package name", func() {
			c := ServiceClient{Name: "MyClient"}

			c.GenerateType(output)
			Expect(output).To(ContainCode(`
				type MyClient struct {
					BaseURL *url.URL
					Client  tigshttp.Client
				}
			`))
		})

		It("should provide a New* function", func() {
			c := ServiceClient{Name: "TestClient"}

			c.GenerateFactoryFunction(output)
			Expect(output).To(ContainCode(`
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
				Expect(client.Imports()).To(ContainElement("net/url"))
				Expect(client.Imports()).To(ContainElement("net/http"))
				Expect(client.Imports()).To(ContainElement("fmt"))
			})

			It("should return all packages necessary if there are json parameters", func() {
				client.Endpoints = []Endpoint{{Method: "POST", Name: "Do", Parameters: []Parameter{
					{Name: "p", Location: "json"},
				}}}

				Expect(client.Imports()).To(ContainElement("bytes"))
				Expect(client.Imports()).To(ContainElement("encoding/json"))
				Expect(client.Imports()).To(ContainElement("io/ioutil"))
				Expect(client.Imports()).To(ContainElement("github.com/fgrosse/tigs/tigshttp"))
			})
		})
	})
})
