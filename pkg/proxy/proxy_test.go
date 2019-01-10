package proxy_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"reverse-proxy/pkg/config"
	"reverse-proxy/pkg/errors"
	. "reverse-proxy/pkg/proxy"
	"reverse-proxy/pkg/respond"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Proxy", func() {
	Context("GET Recipe By ID over reverse proxy", func() {
		CustomURL := "127.0.0.1:5555"
		var ts *httptest.Server

		BeforeEach(func() {
			ts = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					respond.WithFail(w, r, errors.BadRequest(err.Error()))
					return
				}
				recipeID := r.URL.Query().Get("id")
				recipeFilePath := fmt.Sprintf("../test/recipes/stubs/%s.json", recipeID)
				if _, err := os.Stat(recipeFilePath); os.IsNotExist(err) {
					respond.WithFail(w, r, errors.NotFound("recipe not found"))
					return
				}
				recipe, err := ioutil.ReadFile(recipeFilePath)
				if err != nil {
					log.Println(err)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(recipe))
			}))
			l, _ := net.Listen("tcp", CustomURL)
			ts.Listener = l
			ts.Start()
		})
		AfterEach(func() {
			ts.Close()
		})

		It("Invalid target url", func() {
			config.ServerRecipeEndpoint = "http://test.com/Segment%%2815197306101420000%29.ts/%d"
			recipe, err := GetRecipe(1)
			Ω(recipe).Should(BeNil())
			Expect(err.Status).To(Equal(http.StatusInternalServerError))
		})
		It("Get receipe over http", func() {
			config.ServerRecipeEndpoint = ts.URL + "/test-golang-recipes/?id=%d"
			recipe, err := GetRecipe(1)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(recipe.ID).To(Equal(1))
		})
		It("should get 404 - recipe not found", func() {
			config.ServerRecipeEndpoint = ts.URL + "/test-golang-recipes/?id=%d"
			recipe, err := GetRecipe(11)
			Ω(recipe).Should(BeNil())
			Expect(err.Status).To(Equal(http.StatusNotFound))
			Expect(err.Error()).To(Equal("The specified recipe does not exist."))
		})
	})

	Context("Reverse Proxy", func() {
		CustomURL := "127.0.0.1:5555"
		var ts *httptest.Server

		BeforeEach(func() {
			ts = httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					respond.WithFail(w, r, errors.BadRequest(err.Error()))
					return
				}
				recipeID := r.URL.Query().Get("id")
				recipeFilePath := fmt.Sprintf("../test/recipes/stubs/%s.json", recipeID)
				if _, err := os.Stat(recipeFilePath); os.IsNotExist(err) {
					respond.WithFail(w, r, errors.NotFound("recipe not found"))
					return
				}
				recipe, err := ioutil.ReadFile(recipeFilePath)
				if err != nil {
					log.Println(err)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(recipe))
			}))
			l, _ := net.Listen("tcp", CustomURL)
			ts.Listener = l
			ts.Start()
		})
		AfterEach(func() {
			ts.Close()
		})

		It("Get receipe over http", func() {
			config.ProxyServerHost = ts.URL + "/test-golang-recipes/"
			w := httptest.NewRecorder()
			r, err := http.NewRequest("GET", "/test-golang-recipes/1", nil)
			Ω(err).ShouldNot(HaveOccurred())
			ReverseProxy(w, r, "1")
		})
	})
})
