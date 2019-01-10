package respond_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRespond(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Respond Suite")
}
