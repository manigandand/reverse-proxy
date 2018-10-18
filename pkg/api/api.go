package api

import (
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/respond"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// BaseRoutes holds all the registered enpoints
	BaseRoutes *Routes
)

// Routes holds the mux route connection
type Routes struct {
	Root    *mux.Router
	Recipes *mux.Router
}

type apiHandler func(w http.ResponseWriter, r *http.Request) *errors.AppError

// InitAPI initiates all the enpoints
func InitAPI() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = mux.NewRouter()
	BaseRoutes.Root.Handle("/", http.HandlerFunc(indexHandler))
	BaseRoutes.Recipes = BaseRoutes.Root.PathPrefix("/recipes").Subrouter()
	InitRecipe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type Welcome struct {
		Message string `json:"message"`
	}
	respond.With(w, r, http.StatusOK, Welcome{Message: "HelloFresh recipe reverse proxy server."})
}

func (f apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		log.Errorf("Error: %s, StatusCode: %d", err.Error(), err.Status)
		respond.Fail(w, r, err)
	}
}

// LogHandler serves handlerfunc
func LogHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
