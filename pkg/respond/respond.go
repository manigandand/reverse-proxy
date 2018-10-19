package respond

import (
	"compress/gzip"
	"encoding/json"
	"manigandand-golang-test/pkg/config"
	"manigandand-golang-test/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var (
	DEFAULT_PAGE_LIMIT = 10
	decoder            *schema.Decoder
)

// Response holds the handlerfunc response
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// PageResponse holds the paginated handlerfunc response
type PageResponse struct {
	Data interface{} `json:"data"`
	Meta MetaPage    `json:"meta"`
}

// Meta holds the status of the request informations
type Meta struct {
	Status  int    `json:"status_code"`
	Message string `json:"error_message,omitempty"`
}

// MetaPage holds the paginated data inforamtions
type MetaPage struct {
	Meta
	Total    int    `json:"total,omitempty"`
	Count    int    `json:"count,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}

// Page holds the paginate informations
type Page struct {
	Offset int `schema:"offset" url:"offset"`
	Limit  int `schema:"limit" url:"limit"`
}

func init() {
	decoder = schema.NewDecoder()
	decoder.ZeroEmpty(true)
	decoder.IgnoreUnknownKeys(true)
}

// NewPage decodes the pagiante information from the request
func NewPage(r *http.Request) *Page {
	page := &Page{Limit: DEFAULT_PAGE_LIMIT}
	if err := decoder.Decode(page, r.Form); err != nil {
		log.Println(err)
	}
	if page.Limit == 0 || page.Limit > DEFAULT_PAGE_LIMIT {
		page.Limit = DEFAULT_PAGE_LIMIT
	}
	return page
}

// Format customize the http response
func Format(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var res Response
	res.Data = data
	res.Meta = Meta{Status: status}
	With(w, r, status, res)
}

// Fail write the error response
func Fail(w http.ResponseWriter, r *http.Request, e *errors.AppError) {
	var res Response
	res.Meta = Meta{Status: e.Status, Message: e.Message}
	With(w, r, e.Status, res)
}

// With sets the response headers, and write response to client
func With(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)
	if status != http.StatusNoContent {
		if _, err := gz.Write(buf); err != nil {
			log.Error("respond.With.error: ", err)
		}
	}
}

// Paginate returns the response container with paginated inforamtions
func Paginate(w http.ResponseWriter, r *http.Request, data interface{}, p *Page, isEOF bool, count int) {
	res := &PageResponse{Data: data}
	params := r.URL.Query()
	// there exists next page
	if !isEOF {
		nextPage := Page{Limit: p.Limit, Offset: p.Limit + p.Offset}
		params.Set("limit", strconv.Itoa(nextPage.Limit))
		params.Set("offset", strconv.Itoa(nextPage.Offset))
		res.Meta.Next = config.APIHost + r.URL.Path + "?" + params.Encode()
	}
	// there exists previous page
	if p.Offset > 0 {
		var newOff int
		if p.Offset-p.Limit <= 0 {
			newOff = 0
		} else {
			newOff = p.Offset - p.Limit
		}
		prevPage := Page{Limit: p.Limit, Offset: newOff}
		params.Set("limit", strconv.Itoa(prevPage.Limit))
		params.Set("offset", strconv.Itoa(prevPage.Offset))
		res.Meta.Previous = config.APIHost + r.URL.Path + "?" + params.Encode()
	}
	res.Meta.Count = count
	res.Meta.Status = http.StatusOK
	With(w, r, res.Meta.Status, res)
}
