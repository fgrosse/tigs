package main

import (
	. "github.com/fgrosse/gomega-matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"io"
)

var _ = Describe("code generation", func() {
	var output io.Writer
	BeforeEach(func() {
		output = &bytes.Buffer{}
	})

	It("should declare package imports", func() {
		c := client{Name: "TestClient", Package: "tigs_test"}
		Expect(generate(output, c)).To(Succeed())
		for _, i := range c.imports() {
			Expect(output).To(ImportPackage(i))
		}
	})

	It("should include the type definition of the client", func() {
		c := client{Name: "MyClient", Package: "my_package"}

		Expect(generate(output, c)).To(Succeed())
		Expect(output).To(DeclarePackage("my_package"))
		Expect(output).To(ContainCode(`type MyClient struct`))
	})

	It("should include the factory function of the client", func() {
		c := client{Name: "TestClient", Package: "tigs_test"}
		Expect(generate(output, c)).To(Succeed())
		Expect(output).To(ContainCode(`func NewTestClient(baseURL string) (*TestClient, error)`))
	})
})
