package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// 	var response recipe.Recipe
	// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	log.Infof("%+v", response)
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
}

var _ http.RoundTripper = &transport{}

func main() {
	target, err := url.Parse("https://s3-eu-west-1.amazonaws.com/test-golang-recipes/1")
	// target, err := url.Parse("http://test-golang-recipes.s3.amazonaws.com/1")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &transport{http.DefaultTransport}
	http.Handle("/", proxy)
	log.Println("Starting server on port :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
