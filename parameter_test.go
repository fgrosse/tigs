package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parameter", func() {
	It("should figure out the correct go type from its type string", func() {
		testData := map[string]string{
			"string":  "string",
			"text":    "string",
			"int":     "int",
			"integer": "int",
			"float":   "float64",
			"float32": "float32",
			"float64": "float64",
			"bool":    "bool",
			"boolean": "bool",
			"":        "interface{}",
			"error":   "error",
			"Foobar":  "Foobar",
		}

		for actual, expected := range testData {
			p := parameter{Type: actual}
			Expect(p.generatedType()).To(Equal(expected))
		}
	})

	It("should generate code to get its string representation", func() {
		testData := map[string]string{
			"string":  `x`,
			"text":    `x`,
			"int":     `fmt.Sprintf("%d", x)`,
			"integer": `fmt.Sprintf("%d", x)`,
			"float":   `fmt.Sprintf("%f", x)`,
			"float32": `fmt.Sprintf("%f", x)`,
			"float64": `fmt.Sprintf("%f", x)`,
			"bool":    `fmt.Sprintf("%t", x)`,
			"boolean": `fmt.Sprintf("%t", x)`,
			"":        `fmt.Sprintf("%s", x)`,
			"Foobar":  `fmt.Sprintf("%s", x)`,
		}

		for actual, expected := range testData {
			p := parameter{Type: actual, Name: "x"}
			Expect(p.stringCode()).To(Equal(expected))
		}
	})

	Describe("Validate", func() {
		var validParameter = func() parameter {
			return parameter{
				Name:        "foo",
				Description: "This is a test endpoint",
				Type:        "string",
				Location:    "query",
			}
		}

		It("should reject parameters without a Name", func() {
			p := validParameter()

			p.Name = ""
			Expect(p.Validate()).To(MatchError("missing name"))
		})

		It("should reject parameters without a type", func() {
			p := validParameter()

			p.Type = ""
			Expect(p.Validate()).To(MatchError("missing type"))
		})

		It("should reject parameters without a Location", func() {
			p := validParameter()

			p.Location = ""
			Expect(p.Validate()).To(MatchError("missing location"))
		})

		It("should reject parameters with an unknown Location", func() {
			p := validParameter()

			p.Location = "foobar"
			Expect(p.Validate()).To(MatchError(`unknown location "foobar"`))
		})
	})
})
