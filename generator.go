package tigs

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// TODO use interface with only once action for http client:
//     Do(req *Request) (resp *Response, err error)
// TODO generate code comments
// TODO check if generated code compiles
// TODO include go generate comment
// TODO generate tigs comments (with version)

func Generate(w io.Writer, c Client) error {
	// TODO test if client is valid and if not then reject it
	out := &formattableWriter{w}

	generatePackage(out, c)
	generateImports(out, c)
	generateTypeDefinition(out, c)
	generateNewTypeFunction(out, c)

	for _, ep := range c.Endpoints {
		ep.client = c
		generateEndpointFunction(out, ep)
	}

	return nil
}

func generatePackage(out *formattableWriter, c Client) {
	out.printf("package %s\n", c.Package)
}

func generateImports(out *formattableWriter, c Client) {
	imports := []string{`"fmt"`, `"net/url"`}

	if len(c.Endpoints) > 0 {
		imports = append(imports, `"net/http"`)
	}

	sort.Strings(imports)
	out.printf("import (\n\t%s\n)\n", strings.Join(imports, "\n\t"))
}

func generateTypeDefinition(out *formattableWriter, c Client) {
	out.printf(`type %s struct {`, c.Name)
	out.printf(`    BaseURL *url.URL`)
	out.printf(`    Client  *http.Client`)
	out.printf(`}`)
}

func generateNewTypeFunction(out *formattableWriter, c Client) {
	out.printf(``)
	out.printf(`func New%s(baseURL string) (*%s, error) {`, c.Name, c.Name)
	out.printf(`    u, err := url.Parse(baseURL)`)
	out.printf(`    if err != nil {`)
	out.printf(`        return nil, fmt.Errorf("invalid base URL for new %s: %%s", err)`, c.Name)
	out.printf(`    }`)
	out.printf(``)
	out.printf(`    return &FooClient{`)
	out.printf(`        BaseURL: u,`)
	out.printf(`        Client: http.DefaultClient,`)
	out.printf(`    }, nil`)
	out.printf(`}`)
}

func generateEndpointFunction(out *formattableWriter, ep Endpoint) {
	if ep.Method == "GET" && ep.URL == "" && len(ep.Parameters) == 0 {
		out.printf(``)
		out.printf(`func (c *%s) %s() (*http.Response, error) {`, ep.client.Name, ep.Name)
		out.printf(`    return c.Client.Get(c.BaseURL.String())`)
		out.printf(`}`)
		out.printf(``)
		return
	}

	args := []string{}
	for _, p := range ep.Parameters {
		args = append(args, fmt.Sprintf(p.Name)+" "+p.GeneratedType())
	}

	out.printf(``)
	out.printf(`func (c *%s) %s(%s) (*http.Response, error) {`, ep.client.Name, ep.Name, strings.Join(args, ", "))
	out.printf(`    u, err := c.BaseURL.Parse(%q)`, ep.URL) // TODO check what happens if baseURL = foobar/v1/ and ep path = /blup
	out.printf(`    if err != nil {`)
	out.printf(`        return nil, err`)
	out.printf(`    }`)
	out.printf(``)

	for _, p := range ep.Parameters {
		out.printf("\tu.Query().Add(%q, %s)", p.Name, p.StringCode())
	}

	if len(ep.Parameters) > 0 {
		out.printf(``)
	}

	switch ep.Method {
	case "GET":
		out.printf("\treturn c.Client.Get(u.String())")
	default:
		panic("NOT IMPLEMENTED")
	}

	out.printf(`}`)
}
