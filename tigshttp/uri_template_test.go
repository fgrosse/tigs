package tigshttp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/tigs/tigshttp"
)

var _ = Describe("URITemplate", func() {
	It("should expand simple variables", func() {
		u, err := tigshttp.ExpandURITemplate("/foo/{bar}/baz/{id}-{state}", map[string]interface{}{
			"bar": "TEST",
			"id": 42,
			"state": true,
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(u).NotTo(BeNil())
		Expect(u.String()).To(Equal("/foo/TEST/baz/42-true"))
	})

	It("should complain about missing }", func() {
		_, err := tigshttp.ExpandURITemplate("/foo/{bar/blup", map[string]interface{}{})

		Expect(err).To(MatchError("invalid URI template: missing }"))
	})

	It("should ignore unknown variables", func() {
		u, err := tigshttp.ExpandURITemplate("/foo/{bar}/blup", map[string]interface{}{})

		Expect(err).NotTo(HaveOccurred())
		Expect(u).NotTo(BeNil())
		Expect(u.String()).To(Equal("/foo/%7Bbar%7D/blup"))
	})
})
