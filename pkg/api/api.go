package api

import (
	"net/http"
	"reverse-proxy/pkg/config"
	"reverse-proxy/pkg/errors"
	"reverse-proxy/pkg/respond"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// BaseRoutes holds all the registered enpoints
	BaseRoutes *Routes
)

// Routes holds the mux route connection
type Routes struct {
	Root       *mux.Router
	Recipe     *mux.Router
	NeedRecipe *mux.Router
	Recipes    *mux.Router
}

type apiHandler func(w http.ResponseWriter, r *http.Request) *errors.AppError

// InitAPI initiates all the enpoints
func InitAPI() {
	BaseRoutes = &Routes{}
	BaseRoutes.Root = mux.NewRouter()
	BaseRoutes.Root.Handle("/", http.HandlerFunc(indexHandler))
	BaseRoutes.Recipe = BaseRoutes.Root.PathPrefix("/recipe").Subrouter()
	BaseRoutes.NeedRecipe = BaseRoutes.Recipe.PathPrefix("/{recipe-id:[0-9]+}").Subrouter()
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
		respond.WithFail(w, r, err)
	}
}

// RecipeRequired validates the request recipe ids
func RecipeRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			IDs          []int
			IDContainers []string
		)
		if err := r.ParseForm(); err != nil {
			respond.WithFail(w, r, errors.BadRequest(err.Error()))
			return
		}

		ids := r.URL.Query().Get("ids")
		if strings.TrimSpace(ids) != "" {
			IDContainers = strings.Split(ids, ",")
		}

		for _, id := range IDContainers {
			if strings.TrimSpace(id) == "" {
				continue
			}
			i, err := strconv.Atoi(id)
			if err != nil {
				respond.WithFail(w, r, errors.BadRequest("invalid recipe id"))
				return
			}
			IDs = append(IDs, i)
		}
		if len(IDs) > config.MaxRecipesIDs {
			respond.WithFail(w, r, errors.UnprocessableEntity("More than 5 recipe ids"))
			return
		}

		context.Set(r, "ids", IDs)
		next.ServeHTTP(w, r)
		return
	}
	return http.HandlerFunc(fn)
}

// LogHandler serves handlerfunc
func LogHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
