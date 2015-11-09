package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"
)

var _ = Describe("Guzzle Service descriptions", func() {
	It("should unmarshal YAML input", func() {
		yaml := `
			name: TestService
			description: An example client for the amazing TestService
			version: "3.14"
			operations:
				fooOperation:
					summary:  This is an example of an abstract operation
					abstract: true
					httpMethod: GET
					parameters:
						page:
							description: Pagination parameter to request a specific page number.
							type: integer
							location: query
						per_page:
							description: Pagination parameter to request the page size.
							type: integer
							location: query
				DoStuff:
					extends: fooOperation
					summary: Some test endpoint
					uri:     this/is/a/test
					parameters:
						name:
							type: string
							location: query
							required: true
`
		d := newDecoder("guzzle-yaml", strings.NewReader(yaml))
		c := new(client)
		Expect(d.decode(c, settings{Inheritance: true})).To(Succeed())
		Expect(c.Name).To(Equal("TestService"))
		Expect(c.Description).To(Equal("An example client for the amazing TestService"))
		Expect(c.APIVersion).To(Equal("3.14"))
		Expect(c.Endpoints).To(ConsistOf(
			endpoint{
				Name:        "DoStuff",
				Extends:     "fooOperation",
				Description: "Some test endpoint",
				Method:      "GET",
				URI:         "this/is/a/test",
				Parameters: []parameter{
					{Name: "name", Type: "string", Location: "query", Required: true},
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
			endpoint{
				Name:        "fooOperation",
				Abstract:    true,
				Description: "This is an example of an abstract operation",
				Method:      "GET",
				Parameters: []parameter{
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
		))
	})

	It("should unmarshal JSON input", func() {
		json := `
		{
			"name": "TestService",
			"description": "An example client for the amazing TestService",
			"version": "3.14",
			"operations": {
				"fooOperation": {
					"summary": "This is an example of an abstract operation",
					"abstract": true,
					"httpMethod": "GET",
					"parameters": {
						"page": {
							"description": "Pagination parameter to request a specific page number.",
							"type": "integer",
							"location": "query"
						},
						"per_page": {
							"description": "Pagination parameter to request the page size.",
							"type": "integer",
							"location": "query"
						}
					}
				},
				"DoStuff": {
					"extends": "fooOperation",
					"summary": "Some test endpoint",
					"uri": "this/is/a/test",
					"parameters": {
						"name": {
							"type": "string",
							"location": "query",
							"required": true
						}
					}
				}
			}
		}
`
		d := newDecoder("guzzle-json", strings.NewReader(json))
		c := new(client)
		Expect(d.decode(c, settings{Inheritance: true})).To(Succeed())
		Expect(c.Name).To(Equal("TestService"))
		Expect(c.Description).To(Equal("An example client for the amazing TestService"))
		Expect(c.APIVersion).To(Equal("3.14"))
		Expect(c.Endpoints).To(ConsistOf(
			endpoint{
				Name:        "DoStuff",
				Extends:     "fooOperation",
				Description: "Some test endpoint",
				Method:      "GET",
				URI:         "this/is/a/test",
				Parameters: []parameter{
					{Name: "name", Type: "string", Location: "query", Required: true},
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
			endpoint{
				Name:        "fooOperation",
				Abstract:    true,
				Description: "This is an example of an abstract operation",
				Method:      "GET",
				Parameters: []parameter{
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
		))
	})

	It("should return an error if an operation is extending an unknown operation", func() {
		yaml := `
			name: TestService
			operations:
				DoStuff:
					extends: fooOperation
`
		d := newDecoder("guzzle-yaml", strings.NewReader(yaml))
		Expect(d.decode(new(client), settings{Inheritance: true})).To(MatchError(`endpoint "DoStuff" extends an unknown endpoint "fooOperation"`))
	})
})
