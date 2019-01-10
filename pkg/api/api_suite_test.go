package api_test

import (
	"io"
	. "reverse-proxy/pkg/api"
	"reverse-proxy/pkg/config"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const HEADER_CONTENT_TYPE = "Content-Type"

var (
	tServer *httptest.Server
	tClient *Client
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

type Client struct {
	URL        string
	HTTPClient *http.Client
}

func NewClient(URL string) *Client {
	return &Client{URL, &http.Client{}}
}

func (c *Client) DoAPIGet(url string) (*http.Response, error) {
	return c.DoAPICall("GET", c.URL+url, nil)
}

func (c *Client) DoAPICall(method, url string, body io.Reader) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set(HEADER_CONTENT_TYPE, "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

var _ = BeforeSuite(func() {
	Setup()
})

var _ = AfterSuite(func() {
})

func MainSetup() {
	InitAPI()
	testPort := "8001"
	os.Setenv("ENV", config.EnvDevelopment)
	os.Setenv("PORT", testPort)
	os.Setenv("API_HOST", "http://localhost:"+testPort)
	os.Setenv("SEREVR_RECIPE_ENDPOINT", "https://s3-eu-west-1.amazonaws.com/test-golang-recipes/%d")

	if tServer == nil {
		tServer = httptest.NewServer(BaseRoutes.Root)
	}
}

func MainTearDown() {
	// flush all the env
	os.Clearenv()
	if tServer != nil {
		tServer.Close()
	}
}

func Setup() {
	if tClient == nil {
		tClient = NewClient(tServer.URL)
	}
}
