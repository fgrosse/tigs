package tigs

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parameter", func() {
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
			p := Parameter{Type: actual}
			Expect(p.GeneratedType()).To(Equal(expected))
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
			p := Parameter{Type: actual, Name: "x"}
			Expect(p.StringCode()).To(Equal(expected))
		}
	})
})
