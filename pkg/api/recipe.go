package api

import (
	"encoding/json"
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
)

// InitRecipe initiates the recipe enpoints
func InitRecipe() {
	// Get the recipe
	BaseRoutes.Recipes.Handle(
		"", RecipeRequired(apiHandler(getRecipeHandler)),
	).Methods(http.MethodGet)
}

func getRecipeHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ids := context.Get(r, "ids").([]int)

	// get recipe by ids
	if len(ids) > 0 {

	}
	// get all recipes

	log.Info(ids)
	recipe, err := getRecipeByID(1)
	if err != nil {
		return errors.InternalServer(err.Error())
	}

	respond.Format(w, r, http.StatusOK, recipe)
	return nil
}

func getRecipeByIDs(ids []int) {
	for _, id := range ids {
		getRecipeByID(id)
	}
}

func getRecipeByID(id int) (*recipe.Recipe, error) {
	var response recipe.Recipe

	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	url, err := url.Parse("https://s3-eu-west-1.amazonaws.com/test-golang-recipes/1")
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
