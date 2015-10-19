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

			err := Generate(output, client)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(DeclarePackage("my_package"))
			Expect(output).To(ContainCode(`
				type MyClient struct {
					BaseURL *url.URL
					Client  *http.Client
				}
			`))
		})

		It("should provide a New* function", func() {
			err := Generate(output, client)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ImportPackage("net/url"))
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
			client.Endpoints = []Endpoint{{
				Name:        "GetStuff",
				Description: "This is just a very simple endpoint for this test",
				Method:      "GET",
			}}

			err := Generate(output, client)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ImportPackage("net/http"))
			Expect(output).To(ContainCode(`
				func (c *TestClient) GetStuff() (*http.Response, error) {
					return c.Client.Get(c.BaseURL.String())
				}
			`))
		})

		It("should generate GET operations", func() {
			client.Endpoints = []Endpoint{{
				Name:        "GetStuff",
				Description: "This is just a very simple endpoint for this test",
				Method:      "GET",
				URL:         "/stuff",
				Parameters: []Parameter{
					{Name: "s", Type: "string"},
					{Name: "b", Type: "bool"},
					{Name: "q", Type: "boolean"},
					{Name: "i", Type: "int"},
					{Name: "n", Type: "integer"},
					{Name: "i32", Type: "int32"},
					{Name: "i64", Type: "int64"},
					{Name: "f32", Type: "float32"},
					{Name: "f64", Type: "float64"},
					{Name: "foo"}, // type interface{} is implicit
				},
			}}

			err := Generate(output, client)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ImportPackage("net/http"))
			Expect(output).To(ContainCode(`
				func (c *TestClient) GetStuff(s string, b bool, q bool, i int, n int, i32 int32, i64 int64, f32 float32, f64 float64, foo interface{}) (*http.Response, error) {
					u, err := c.BaseURL.Parse("/stuff")
					if err != nil {
						return nil, err
					}

					u.Query().Add("s", s)
					u.Query().Add("b", fmt.Sprintf("%t", b))
					u.Query().Add("q", fmt.Sprintf("%t", q))
					u.Query().Add("i", fmt.Sprintf("%d", i))
					u.Query().Add("n", fmt.Sprintf("%d", n))
					u.Query().Add("i32", fmt.Sprintf("%d", i32))
					u.Query().Add("i64", fmt.Sprintf("%d", i64))
					u.Query().Add("f32", fmt.Sprintf("%f", f32))
					u.Query().Add("f64", fmt.Sprintf("%f", f64))
					u.Query().Add("foo", fmt.Sprintf("%s", foo))

					return c.Client.Get(u.String())
				}
			`))
		})

		PIt("should generate POST operations", func() {
			client.Endpoints = []Endpoint{{
				Name:        "GetStuff",
				Description: "This is just a very simple endpoint for this test",
				Method:      "GET",
				Parameters: []Parameter{
					{Name: "s", Type: "string", Location: "query"},
					{Name: "b", Type: "bool", Location: "json"},
					{Name: "i", Type: "int", Location: "json"},
					{Name: "foo", Location: "json"},
				},
			}}

			err := Generate(output, client)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(ImportPackage("net/http"))
			Expect(output).To(ContainCode(`
				func (c *TestClient) GetStuff(s string, b bool, i int, foo interface{}) (*http.Response, error) {
					u, err := c.BaseURL.Parse("/stuff")
					if err != nil {
						return nil, err
					}

					u.Query().Add("s", s)

					data, err := json.Marshal(map[string]interface{}{
						"b":   b,
						"i":   i,
						"foo": foo,
					})

					if err != nil {
						return nil, fmt.Errorf("could not marshal body for PostStuff: %s", err)
					}

					return c.Client.Post(u.String(), "application/json", bytes.NewBuffer(data))
				}
			`))
		})
	})
})
