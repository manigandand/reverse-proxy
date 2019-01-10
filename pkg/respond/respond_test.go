package respond_test

import (
	"net/http"
	"net/http/httptest"
	"reverse-proxy/pkg/errors"
	. "reverse-proxy/pkg/respond"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTestRequest() *http.Request {
	r, err := http.NewRequest("GET", "testing", nil)
	if err != nil {
		panic("request create failed: " + err.Error())
	}
	return r
}

var _ = Describe("Respond", func() {
	var testData = map[string]interface{}{"test": true}
	It("Test With - should respond custom data", func() {
		w := httptest.NewRecorder()
		r := newTestRequest()
		With(w, r, http.StatusOK, testData)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})
	It("New Page - should decode paginate information from request", func() {
		r, err := http.NewRequest("GET", "recipe", nil)
		Ω(err).ShouldNot(HaveOccurred())
		urls := r.URL.Query()
		urls.Add("limit", "10")
		urls.Add("offset", "10")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()

		page := NewPage(r)
		Expect(page.Limit).To(Equal(10))
		Expect(page.Offset).To(Equal(10))
	})

	It("New Page - max limit should be 10. server should reset limit to 10", func() {
		r, err := http.NewRequest("GET", "recipe", nil)
		urls := r.URL.Query()
		urls.Add("limit", "20")
		urls.Add("offset", "10")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()

		Ω(err).ShouldNot(HaveOccurred())
		page := NewPage(r)
		Expect(page.Limit).To(Equal(10))
		Expect(page.Offset).To(Equal(10))
	})
	It("Foramt - should respond formated response", func() {
		w := httptest.NewRecorder()
		r := newTestRequest()
		Format(w, r, http.StatusOK, testData)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})
	It("WithFail - should return error response", func() {
		w := httptest.NewRecorder()
		r := newTestRequest()
		WithFail(w, r, errors.BadRequest("invalid recipe id"))
		Expect(w.Code).To(Equal(http.StatusBadRequest))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})

	It("Paginate - should return paginated response", func() {
		w := httptest.NewRecorder()
		r := newTestRequest()
		r, err := http.NewRequest("GET", "recipe", nil)
		Ω(err).ShouldNot(HaveOccurred())
		urls := r.URL.Query()
		urls.Add("limit", "10")
		urls.Add("offset", "10")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()
		page := NewPage(r)

		Paginate(w, r, testData, page, false, 10)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})

	It("Paginate - should return paginated response with previous link", func() {
		w := httptest.NewRecorder()
		r := newTestRequest()
		r, err := http.NewRequest("GET", "recipe", nil)
		Ω(err).ShouldNot(HaveOccurred())
		urls := r.URL.Query()
		urls.Add("limit", "10")
		urls.Add("offset", "50")
		r.URL.RawQuery = urls.Encode()
		r.ParseForm()
		page := NewPage(r)

		Paginate(w, r, testData, page, false, 10)
		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
		Expect(w.HeaderMap.Get("Content-Encoding")).To(Equal("gzip"))
	})
})
