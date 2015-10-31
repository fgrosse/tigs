package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func generate(w io.Writer, c client) error {
	out := &formattableWriter{w}
	out.printf("package %s\n", c.Package)

	imports := c.imports()
	for i, s := range imports {
		imports[i] = `"` + s + `"`
	}

	sort.Strings(imports)
	out.printf("import (\n\t%s\n)\n", strings.Join(imports, "\n\t"))

	if len(c.Name) < 6 || c.Name[len(c.Name)-6:] != "Client" {
		c.Name = c.Name + "Client"
	}

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
