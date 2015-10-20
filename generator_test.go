package tigs

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
		client Client
	)

	BeforeEach(func() {
		output = &bytes.Buffer{}
		client = Client{Name: "TestClient", Package: "tigs_test"}
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
					Client  *http.Client
				}
			`))
		})

		It("should provide a New* function", func() {
			Expect(Generate(output, client)).To(Succeed())
			Expect(output).To(ImportPackage("net/url"))
			Expect(output).To(ImportPackage("net/http"))
			Expect(output).To(ContainCode(`
				func NewTestClient(baseURL string) (*TestClient, error) {
					u, err := url.Parse(baseURL)
					if err != nil {
						return nil, fmt.Errorf("invalid base URL for new TestClient: %s", err)
					}

					return &FooClient{
						BaseURL: u,
						Client: http.DefaultClient,
					}, nil
				}
			`))
		})

		It("should generate simple GET operations without arguments", func() {
			client.Endpoints = []Endpoint{{Name: "GetStuff", Method: "GET"}}

			Expect(Generate(output, client)).To(Succeed())
			Expect(output).To(ContainCode(`
				func (c *TestClient) GetStuff() (*http.Response, error) {
					return c.Client.Get(c.BaseURL.String())
				}
			`))
		})

		It("should generate GET operations", func() {
			client.Endpoints = []Endpoint{{
				Name:   "GetStuff",
				Method: "GET", URL: "/stuff",
				Parameters: []Parameter{
					{Name: "s", Type: "string"},
					{Name: "i", Type: "int"},
				},
			}}

			Expect(Generate(output, client)).To(Succeed())
			Expect(output).To(ImportPackage("net/http"))
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
			client.Endpoints = []Endpoint{{
				Name:   "CreateStuff",
				Method: "POST", URL: "/stuff",
				Parameters: []Parameter{
					{Name: "s", Type: "string", Location: "query"},
					{Name: "b", Type: "bool", Location: "json"},
					{Name: "i", Type: "int", Location: "json"},
				},
			}}

			Expect(Generate(output, client)).To(Succeed())
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
})
