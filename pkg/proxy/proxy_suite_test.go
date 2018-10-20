package proxy_test

import (
	"manigandand-golang-test/pkg/config"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Proxy Suite")
}

var _ = AfterSuite(func() {
	config.ServerRecipeEndpoint = ""
	config.ProxyServerHost = ""
	// flush all the env
	os.Clearenv()
})
