package respond

import (
	"compress/gzip"
	"encoding/json"
	"manigandand-golang-test/pkg/errors"
	"net/http"

	log "github.com/sirupsen/logrus"
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
	Total    int    `json:"total"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)
	if status != http.StatusNoContent {
		if _, err := gz.Write(buf); err != nil {
			log.Error("respond.With.error: ", err)
		}
	}
}
