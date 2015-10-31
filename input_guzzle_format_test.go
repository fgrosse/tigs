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
	paginatedOperation:
		summary:  This is an example of an abstract operation
		abstract: true
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
		summary: Some test endpoint
		method:  GET
		uri:     this/is/a/test
		parameters:
			name:
				type: string
				location: query
				required: true
`
		d := newDecoder("guzzle-yaml", strings.NewReader(yaml))
		c := new(client)
		Expect(d.decode(c)).To(Succeed())
		Expect(c.Name).To(Equal("TestService"))
		Expect(c.Description).To(Equal("An example client for the amazing TestService"))
		Expect(c.APIVersion).To(Equal("3.14"))
		Expect(c.Endpoints).To(ConsistOf([]endpoint{
			{
				Name:        "DoStuff",
				Description: "Some test endpoint",
				Method:      "GET",
				URL:         "this/is/a/test",
				Parameters: []parameter{
					{Name: "name", Type: "string", Location: "query", Required: true},
				},
			},
			{
				Name:        "paginatedOperation",
				Abstract:    true,
				Description: "This is an example of an abstract operation",
				Parameters: []parameter{
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
		}))
	})

	It("should unmarshal JSON input", func() {
		json := `
		{
			"name": "TestService",
			"description": "An example client for the amazing TestService",
			"version": "3.14",
			"operations": {
				"paginatedOperation": {
					"summary": "This is an example of an abstract operation",
					"abstract": true,
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
					"summary": "Some test endpoint",
					"method": "GET",
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
		Expect(d.decode(c)).To(Succeed())
		Expect(c.Name).To(Equal("TestService"))
		Expect(c.Description).To(Equal("An example client for the amazing TestService"))
		Expect(c.APIVersion).To(Equal("3.14"))
		Expect(c.Endpoints).To(ConsistOf([]endpoint{
			{
				Name:        "DoStuff",
				Description: "Some test endpoint",
				Method:      "GET",
				URL:         "this/is/a/test",
				Parameters: []parameter{
					{Name: "name", Type: "string", Location: "query", Required: true},
				},
			},
			{
				Name:        "paginatedOperation",
				Abstract:    true,
				Description: "This is an example of an abstract operation",
				Parameters: []parameter{
					{Name: "page", Type: "integer", Location: "query", Description: "Pagination parameter to request a specific page number."},
					{Name: "per_page", Type: "integer", Location: "query", Description: "Pagination parameter to request the page size."},
				},
			},
		}))
	})
})
