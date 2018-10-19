package proxy

import (
	"encoding/json"
	"fmt"
	"manigandand-golang-test/pkg/config"
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// GetRecipe get the recipe details from server over HTTP
func GetRecipe(recipeID int) (*recipe.Recipe, *errors.AppError) {
	var response recipe.Recipe

	client := &http.Client{
		Timeout: config.ClientTimeout,
	}
	url, err := url.Parse(fmt.Sprintf(config.ServerRecipeEndpoint, recipeID))
	if err != nil {
		return nil, errors.InternalServer(err.Error())
	}
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, errors.InternalServer(err.Error())
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.InternalServer(err.Error())
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.NotFound("The specified recipe does not exist.")
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.InternalServer(err.Error())
	}

	return &response, nil
}

// ReverseProxy Serve a reverse proxy for a given url
func ReverseProxy(w http.ResponseWriter, req *http.Request, recipeID int) {
	url, _ := url.Parse(fmt.Sprintf(config.ServerRecipeEndpoint, recipeID))
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Scheme = url.Scheme
	req.URL.Host = url.Host
	req.URL.Path = url.Path
	req.Host = url.Host
	proxy.ServeHTTP(w, req)
}
