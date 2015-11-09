package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("endpointTree", func() {
	It("should copy all parameters from extended endpoints", func() {
		eps := []endpoint{
			{
				Name:     "Foo",
				Abstract: true,
				Parameters: []parameter{
					{Name: "p1", Type: "string"},
					{Name: "p2", Type: "int"},
				},
			},
			{
				Name:    "Bar",
				Extends: "Foo",
				Parameters: []parameter{
					{Name: "p3", Type: "float"},
					{Name: "p4", Type: "bool"},
				},
			},
		}

		Expect(newEndpointTree(eps).process()).To(Succeed())
		Expect(eps[1].Parameters).To(ConsistOf([]parameter{
			{Name: "p1", Type: "string"},
			{Name: "p2", Type: "int"},
			{Name: "p3", Type: "float"},
			{Name: "p4", Type: "bool"},
		}))
	})

	It("should copy method and url from parents", func() {
		eps := []endpoint{
			{
				Name:   "Foo",
				Method: "GET",
				URI:    "/foo",
			},
			{
				Name:    "Bar",
				Extends: "Foo",
			},
			{
				Name:    "Baz",
				Extends: "Foo",
				Method:  "POST",
			},
			{
				Name:    "Blup",
				Extends: "Foo",
				Method:  "POST",
				URI:     "/blup",
			},
		}

		Expect(newEndpointTree(eps).process()).To(Succeed())
		Expect(eps[1].Method).To(Equal("GET"))
		Expect(eps[1].URI).To(Equal("/foo"))
		Expect(eps[2].Method).To(Equal("POST"))
		Expect(eps[2].URI).To(Equal("/foo"))
		Expect(eps[3].Method).To(Equal("POST"))
		Expect(eps[3].URI).To(Equal("/blup"))
	})

	It("should work over multiple levels of inheritance", func() {
		eps := []endpoint{
			{
				Name:       "Foo",
				Method:     "GET",
				URI:        "/foo",
				Parameters: []parameter{{Name: "p1", Type: "string"}},
			},
			{
				Name:       "Bar",
				Extends:    "Foo",
				URI:        "/bar",
				Parameters: []parameter{{Name: "p2", Type: "int"}},
			},
			{
				Name:       "Baz",
				Extends:    "Bar",
				Parameters: []parameter{{Name: "p3", Type: "float"}},
			},
			{
				Name:       "Blup",
				Extends:    "Baz",
				Parameters: []parameter{{Name: "p4", Type: "bool"}},
			},
		}

		Expect(newEndpointTree(eps).process()).To(Succeed())
		Expect(eps[0].Method).To(Equal("GET"))
		Expect(eps[0].URI).To(Equal("/foo"))
		Expect(eps[0].Parameters).To(ConsistOf([]parameter{
			{Name: "p1", Type: "string"},
		}))
		Expect(eps[1].Method).To(Equal("GET"))
		Expect(eps[1].URI).To(Equal("/bar"))
		Expect(eps[1].Parameters).To(ConsistOf([]parameter{
			{Name: "p1", Type: "string"},
			{Name: "p2", Type: "int"},
		}))
		Expect(eps[1].Method).To(Equal("GET"))
		Expect(eps[1].URI).To(Equal("/bar"))
		Expect(eps[2].Parameters).To(ConsistOf([]parameter{
			{Name: "p1", Type: "string"},
			{Name: "p2", Type: "int"},
			{Name: "p3", Type: "float"},
		}))
		Expect(eps[1].Method).To(Equal("GET"))
		Expect(eps[1].URI).To(Equal("/bar"))
		Expect(eps[3].Parameters).To(ConsistOf([]parameter{
			{Name: "p1", Type: "string"},
			{Name: "p2", Type: "int"},
			{Name: "p3", Type: "float"},
			{Name: "p4", Type: "bool"},
		}))
	})

	It("should detect dependency cycles and return an error", func() {
		eps := []endpoint{
			{Name: "Foo", Extends: "Bar"},
			{Name: "Bar", Extends: "Baz"},
			{Name: "Baz", Extends: "Foo"},
		}

		Expect(newEndpointTree(eps).process()).To(
			MatchError("Detected inheritance cycle: Foo -> Bar -> Baz -> Foo"),
		)
	})
})
