package api_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	MainSetup()
	defer MainTearDown()
	os.Exit(m.Run())
}

var _ = Describe("Api", func() {
	Context("GET Index http://127.0.0.1:8080/", func() {
		It("GET crawler index", func() {
			res, err := tClient.ProxyRecipeHomePage()
			Ω(err).ShouldNot(HaveOccurred())
			var data struct {
				Message string `json:"message"`
			}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			// fmt.Println(data)
			Expect(data.Message).To(Equal("HelloFresh recipe reverse proxy server."))
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})
	})
})

func (c *Client) ProxyRecipeHomePage() (*http.Response, error) {
	return c.DoAPIGet("/")
}
