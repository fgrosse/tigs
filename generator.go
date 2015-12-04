package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/fgrosse/gotility"
)

func generate(w io.Writer, c *client) error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("invalid client: %s", err)
	}

	template := loadTemplate("templates/client.tmpl")

	generateTypeName(c)
	generateImports(c)
	generateTypeComment(c)

	return template.Execute(w, c)
}

func generateTypeName(c *client) {
	if len(c.Name) < 6 || c.Name[len(c.Name)-6:] != "Client" {
		c.Name = c.Name + "Client"
	}
}

func generateImports(c *client) {
	stdImportsSet := gotility.NewStringSet(`"fmt"`, `"net/http"`, `"net/url"`)
	otherImports := []string{`"github.com/fgrosse/tigs/tigshttp"`}

	if c.containsJSONEndpoints() {
		stdImportsSet.Set(`"encoding/json"`)
		stdImportsSet.Set(`"bytes"`)
		stdImportsSet.Set(`"io/ioutil"`)
	}

	if c.containsPostfieldEndpoints() {
		stdImportsSet.Set(`"strings"`)
		stdImportsSet.Set(`"io/ioutil"`)
	}

	stdImports := stdImportsSet.All()

	sort.Strings(stdImports)
	sort.Strings(otherImports)
	c.Imports = strings.Join(stdImports, "\n\t") + "\n\n\t" + strings.Join(otherImports, "\n\t")
	for i := range c.Endpoints {
		c.Endpoints[i].ClientName = c.Name
	}
}

func generateTypeComment(c *client) {
	if c.Description == "" {
		c.Description = c.Name + " is an automatically generated HTTP client."
	}

	if len(c.Description) < len(c.Name) || c.Description[:len(c.Name)] != c.Name {
		c.Description = c.Name + ": " + c.Description
	}

	c.Description = strings.TrimSpace(c.Description)
	c.Description = strings.Replace(c.Description, "\n", "\n// ", -1)
}
