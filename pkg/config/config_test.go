package config_test

import (
	. "manigandand-golang-test/pkg/config"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("Test MustEnv", func() {
		BeforeEach(func() {
			os.Setenv("ENV", EnvDevelopment)
			os.Setenv("PORT", "8001")
			os.Setenv("API_HOST", DefaultAPIHost)
			os.Setenv("SEREVR_RECIPE_ENDPOINT", "https://my.bucket.com/recipe/")
		})
		AfterEach(func() {
			// flush all the env
			os.Clearenv()
		})

		It("Should read env values and assign to config variables", func() {
			GetAllEnv()
			Expect(Env).To(Equal(EnvDevelopment))
			Expect(Port).To(Equal("8001"))
			Expect(APIHost).To(Equal(DefaultAPIHost))
			Expect(ServerRecipeEndpoint).To(Equal("https://my.bucket.com/recipe/"))
		})
		It("Should update the env", func() {
			os.Setenv("SEREVR_RECIPE_ENDPOINT", "https://my.new-bucket.com/recipe/")
			GetAllEnv()
			Expect(ServerRecipeEndpoint).To(Equal("https://my.new-bucket.com/recipe/"))
		})
	})
})
