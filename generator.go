package main

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

func Generate(w io.Writer, c ServiceClient) error {
	// TODO test if client is valid and if not then reject it
	out := &formattableWriter{w}
	out.printf("package %s\n", c.Package)

	imports := c.Imports()
	for i, s := range imports {
		imports[i] = `"`+s+`"`
	}

	sort.Strings(imports)
	out.printf("import (\n\t%s\n)\n", strings.Join(imports, "\n\t"))

	c.GenerateType(out)
	c.GenerateFactoryFunction(out)

	for _, ep := range c.Endpoints {
		ep.Generate(w, c.Name)
	}

	return nil
}

