package tigs

import (
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
		ep.Generate(w, c.Name)
	}

	return nil
}

func generatePackage(out *formattableWriter, c Client) {
	out.printf("package %s\n", c.Package)
}

func generateImports(out *formattableWriter, c Client) {
	imports := []string{
		`"fmt"`,
		`"net/http"`,
		`"net/url"`,
	}

	if c.ContainsJSONEndpoints() {
		imports = append(imports, `"encoding/json"`)
		imports = append(imports, `"bytes"`)
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
	out.printf(`    return &%s{`, c.Name)
	out.printf(`        BaseURL: u,`)
	out.printf(`        Client: http.DefaultClient,`)
	out.printf(`    }, nil`)
	out.printf(`}`)
}
