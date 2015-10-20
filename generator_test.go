package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"io"
)

var _ = Describe("Code generation test", func() {
	var (
		output io.Writer
		client ServiceClient
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
		client = ServiceClient{Name: "TestClient", Package: "tigs_test"}
	})

	Describe("generating code", func() {
		It("should define the client and use the given package name", func() {
			client.Name = "MyClient"
			client.Package = "my_package"

			Expect(Generate(output, client)).To(Succeed())
			Expect(output).To(DeclarePackage("my_package"))
			Expect(output).To(ContainCode(`
				type MyClient struct {
					BaseURL *url.URL
					Client  tigshttp.Client
				}
			`))
		})

		It("should provide a New* function", func() {
			Expect(Generate(output, client)).To(Succeed())
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

		Describe("package imports", func() {
			It("should import all packages for the New* function", func() {
				Expect(Generate(output, client)).To(Succeed())
				Expect(output).To(ImportPackage("net/url"))
				Expect(output).To(ImportPackage("net/http"))
				Expect(output).To(ImportPackage("fmt"))
			})

			It("should import the json and bytes package if necessary", func() {
				client.Endpoints = []Endpoint{{Method: "POST", Name: "Do", Parameters: []Parameter{
					{Name: "p", Location: "json"},
				}}}

				Expect(Generate(output, client)).To(Succeed())
				Expect(output).To(ImportPackage("bytes"))
				Expect(output).To(ImportPackage("encoding/json"))
				Expect(output).To(ImportPackage("io/ioutil"))
				Expect(output).To(ImportPackage("github.com/fgrosse/tigs/tigshttp"))
			})
		})
	})
})
