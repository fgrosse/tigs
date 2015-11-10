package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client", func() {
	Describe("Validate", func() {
		It("should reject clients without a Package", func() {
			c := validClient()

			c.Package = ""
			Expect(c.Validate()).To(MatchError("missing package"))
		})

		It("should reject clients without a single endpoint", func() {
			c := validClient()

			c.Endpoints = []endpoint{}
			Expect(c.Validate()).To(MatchError("no endpoints"))

			c.Endpoints = nil
			Expect(c.Validate()).To(MatchError("no endpoints"))
		})

		It("should reject clients with invalid endpoints", func() {
			c := validClient()

			c.Endpoints[0].Name = ""
			Expect(c.Endpoints[0].Validate()).NotTo(Succeed())
			Expect(c.Validate()).To(MatchError(MatchRegexp(`invalid endpoint "": .+`)))
		})
	})
})

func validClient() *client {
	return &client{
		Package:     "MyPackage",
		Name:        "MyClient",
		Description: "This is a test client",
		APIVersion:  "1.2.3",
		Endpoints: []endpoint{
			{Name: "DoFoo", Method: "GET", URI: "/foo"},
		},
	}
}
