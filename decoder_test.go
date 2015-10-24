package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"
)

var _ = Describe("decoder", func() {
	It("should decode from YAML input", func() {
		yaml := `
name: TestService
version: "3.14"
operations:
	DoStuff:
		summary: Some test endpoint
		method:  GET
		uri:     this/is/a/test
		parameters:
			page:
				description: The requested page number
				type: integer
				location: query
			name:
				type: string
				location: query
`
		d := newYAMLDecoder(strings.NewReader(yaml))
		c := new(client)
		Expect(d.Decode(c)).To(Succeed())
		Expect(c.Name).To(Equal("TestService"))
		Expect(c.APIVersion).To(Equal("3.14"))
		Expect(c.Endpoints).To(ConsistOf([]endpoint{
			{
				Name:        "DoStuff",
				Description: "Some test endpoint",
				Method:      "GET",
				URL:         "this/is/a/test",
				Parameters:  []parameter{
					{Name: "name", TypeString: "string", Location: "query"},
					{Name: "page", TypeString: "integer", Location: "query", Description: "The requested page number"},
				},
			},
		}))
	})
})
