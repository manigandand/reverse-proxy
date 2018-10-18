package main

import (
	"encoding/json"
	"manigandand-golang-test/schemas"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var response schemas.Recipe
	// proxyStr := "http://195.201.249.128:3128"
	// proxyURL, err := url.Parse(proxyStr)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	client := &http.Client{
		// Transport: &http.Transport{
		// 	Proxy: http.ProxyURL(proxyURL),
		// },
		Timeout: 2 * time.Second,
	}

	url, err := url.Parse("https://s3-eu-west-1.amazonaws.com/test-golang-recipes/1")
	if err != nil {
		log.Fatal(err.Error())
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("%+v", response)
}
