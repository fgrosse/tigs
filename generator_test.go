package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"io"
)

var _ = Describe("generator", func() {
	var output io.Writer
	BeforeEach(func() {
		output = &bytes.Buffer{}
	})

	Describe("package import", func() {
		It("should return all packages necessary for the factory function", func() {
			c := client{Package: "test_package"}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(ImportPackage("net/url"))
			Expect(output).To(ImportPackage("net/http"))
			Expect(output).To(ImportPackage("fmt"))
		})

		It("should return all packages necessary if there are json parameters", func() {
			c := client{
				Package: "test_package",
				Endpoints: []endpoint{
					{Method: "POST", Name: "Do", Parameters: []parameter{{Name: "p", Location: "json"}}},
				},
			}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(ImportPackage("bytes"))
			Expect(output).To(ImportPackage("encoding/json"))
			Expect(output).To(ImportPackage("io/ioutil"))
			Expect(output).To(ImportPackage("github.com/fgrosse/tigs/tigshttp"))
		})
	})

	Describe("type definition", func() {
		It("should define the client and the correct package", func() {
			c := client{Name: "MyClient", Package: "my_package"}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(DeclarePackage("my_package"))
			Expect(output).To(ContainCode(`
			type MyClient struct {
				BaseURL *url.URL
				Client  tigshttp.Client
			}
		`))
		})

		It("should generate a type comment", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "MyClient is awesome"}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(ContainCode(`
				// MyClient is awesome
				type MyClient struct {
			`))
		})

		It("should always start type comments with the type name", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "This is some description"}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(ContainCode(`
				// MyClient: This is some description
				type MyClient struct {
			`))
		})

		It("should prepend each new line with `//` to mark it as comment", func() {
			c := client{Name: "MyClient", Package: "my_package", Description: "This is some description\nover multiple lines"}

			Expect(generate(output, c)).To(Succeed())
			Expect(output).To(ContainCode(`
				// MyClient: This is some description
				// over multiple lines
				type MyClient struct {
			`))
		})
	})

	It("should provide a New* function", func() {
		c := client{Name: "TestClient", Package: "tigs_test"}

		Expect(generate(output, c)).To(Succeed())
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

	It("should include the factory function of the client", func() {
		c := client{Name: "TestClient", Package: "tigs_test"}
		Expect(generate(output, c)).To(Succeed())
		Expect(output).To(ContainCode(`func NewTestClient(baseURL string) (*TestClient, error)`))
	})

	It("should add `Client` to the type name", func() {
		c := client{Name: "Foo", Package: "my_package"}

		Expect(generate(output, c)).To(Succeed())
		Expect(output).To(ContainCode(`type FooClient struct`))
		Expect(output).To(ContainCode(`func NewFooClient(baseURL string) (*FooClient, error)`))
	})
})
