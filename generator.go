package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// TODO generate code comments
// TODO check if generated code compiles
// TODO include go generate comment
// TODO generate tigs comments (with version)

func generate(w io.Writer, c client) error {
	// TODO test if client is valid and if not then reject it
	out := &formattableWriter{w}
	out.printf("package %s\n", c.Package)

	imports := c.imports()
	for i, s := range imports {
		imports[i] = `"` + s + `"`
	}

	sort.Strings(imports)
	out.printf("import (\n\t%s\n)\n", strings.Join(imports, "\n\t"))

	c.generateType(out)
	c.generateFactoryFunction(out)

	for _, ep := range c.Endpoints {
		ep.generate(out, c.Name)
	}

	return nil
}

type formattableWriter struct {
	io.Writer
}

func (w *formattableWriter) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}
