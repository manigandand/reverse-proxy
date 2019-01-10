package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"manigandand-golang-test/pkg/config"
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api Endpoints", func() {
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
		It("GET /recipe/{id} - should get recipe details", func() {
			config.ProxyServerHost = ts.URL + "/test-golang-recipes/"
			res, err := tClient.ProxyGetRecipeByID(1)
			fmt.Println(res.StatusCode)
			Ω(err).ShouldNot(HaveOccurred())
		})
		It("GET /recipe/{id} - should return 404", func() {
			config.ProxyServerHost = ts.URL + "/test-golang-recipes/"
			res, err := tClient.ProxyGetRecipeByID(11)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			Expect(res.StatusCode).To(Equal(http.StatusNotFound))
		})
	})
	Context("GET Recipes by ids", func() {
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
		It("GET /recipes?ids=1,2,3 - should return 422", func() {
			res, err := tClient.ProxyGetRecipesByID("1,2,3,4,5,6,7,")
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusUnprocessableEntity))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			var data respond.Response
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			Expect(data.Meta.Status).To(Equal(http.StatusUnprocessableEntity))
			Expect(data.Meta.Message).To(Equal("More than 5 recipe ids"))
		})
		It("GET /recipes?ids=1,2,3 - should return 400", func() {
			res, err := tClient.ProxyGetRecipesByID("1,x,x1m")
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			var data respond.Response
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			Expect(data.Meta.Status).To(Equal(http.StatusBadRequest))
			Expect(data.Meta.Message).To(Equal("invalid recipe id"))
		})
		It("GET /recipes?ids=1,2,3 - should return 200, of two recipes", func() {
			config.ServerRecipeEndpoint = ts.URL + "/test-golang-recipes/?id=%d"
			res, err := tClient.ProxyGetRecipesByID("1,2")
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			var data struct {
				Data []recipe.Recipe `json:"data"`
				Meta respond.Meta    `json:"meta"`
			}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			Expect(len(data.Data)).To(Equal(2))
			Expect(data.Data[0].ID).To(Equal(1))
			Expect(data.Data[1].ID).To(Equal(2))
			Expect(data.Meta.Status).To(Equal(http.StatusOK))
		})
	})
	Context("GET All Recipes", func() {
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
		It("GET /recipes - should paginated recipes", func() {
			config.ServerRecipeEndpoint = ts.URL + "/test-golang-recipes/?id=%d"
			res, err := tClient.ProxyGetRecipes(3, 0)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			var data struct {
				Data []recipe.Recipe  `json:"data"`
				Meta respond.MetaPage `json:"meta"`
			}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			Expect(len(data.Data)).To(Equal(3))
			Expect(data.Meta.Status).To(Equal(http.StatusOK))
			Expect(data.Meta.Count).To(Equal(3))
			Expect(data.Meta.Next).To(Equal("http://localhost:8080/recipes?limit=3&offset=3"))
		})
		It("GET /recipes - should paginated recipes", func() {
			config.ServerRecipeEndpoint = ts.URL + "/test-golang-recipes/?id=%d"
			res, err := tClient.ProxyGetRecipes(3, 5)
			Ω(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
			Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
			var data struct {
				Data []recipe.Recipe  `json:"data"`
				Meta respond.MetaPage `json:"meta"`
			}
			jsonErr := json.NewDecoder(res.Body).Decode(&data)
			// fmt.Printf("%+v", data)
			Ω(jsonErr).ShouldNot(HaveOccurred())
			Expect(len(data.Data)).To(Equal(3))
			Expect(data.Meta.Status).To(Equal(http.StatusOK))
			Expect(data.Meta.Count).To(Equal(3))
			Expect(data.Meta.Previous).To(Equal("http://localhost:8080/recipes?limit=3&offset=2"))
			Expect(data.Meta.Next).To(Equal("http://localhost:8080/recipes?limit=3&offset=8"))
		})
	})
})

func (c *Client) ProxyGetRecipeByID(recipeID int) (*http.Response, error) {
	targetPath := fmt.Sprintf("/recipe/%d", recipeID)
	return c.DoAPIGet(targetPath)
}

func (c *Client) ProxyGetRecipesByID(ids string) (*http.Response, error) {
	targetPath := fmt.Sprintf("/recipes?ids=%s", ids)
	return c.DoAPIGet(targetPath)
}

func (c *Client) ProxyGetRecipes(limit, offset int) (*http.Response, error) {
	targetPath := fmt.Sprintf("/recipes?limit=%d&offset=%d", limit, offset)
	return c.DoAPIGet(targetPath)
}
