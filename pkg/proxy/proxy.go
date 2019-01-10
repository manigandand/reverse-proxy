package proxy

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverse-proxy/pkg/config"
	"reverse-proxy/pkg/errors"
	"reverse-proxy/pkg/recipe"
	"strconv"
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
func ReverseProxy(w http.ResponseWriter, r *http.Request, path string) {
	remote, _ := url.Parse(config.ProxyServerHost)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Transport = &transport{http.DefaultTransport}
	proxy.ModifyResponse = modifyProxyResponse

	r.URL.Scheme = remote.Scheme
	r.URL.Host = remote.Host
	r.URL.Path = path
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)

	return
}

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp, nil
}

func modifyProxyResponse(resp *http.Response) (err error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	defer gz.Close()
	_, err = gz.Write(b)
	if err != nil {
		return err
	}
	if err = gz.Flush(); err != nil {
		return err
	}
	compressedData := buf.Bytes()
	resp.Body = ioutil.NopCloser(bytes.NewReader(compressedData))
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("Content-Encoding", "gzip")

	return nil
}
