package tigshttp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTigshttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tigshttp Suite")
}
