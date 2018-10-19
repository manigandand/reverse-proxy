package api

import (
	"encoding/json"
	"fmt"
	"manigandand-golang-test/pkg/config"
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// InitRecipe initiates the recipe enpoints
func InitRecipe() {
	// Get the recipes
	BaseRoutes.NeedRecipe.Handle(
		"", RecipeRequired(apiHandler(getRecipeHandler)),
	).Methods(http.MethodGet)
	// Get the recipes
	BaseRoutes.Recipes.Handle(
		"", RecipeRequired(apiHandler(getRecipesHandler)),
	).Methods(http.MethodGet)
}

// getRecipeHandler
func getRecipeHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	recipeID := getID(r, "recipe-id")
	recipe, err := getRecipeByID(recipeID)
	if err != nil {
		return err
	}

	respond.Format(w, r, http.StatusOK, recipe)
	return nil
}

func getRecipesHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	recipeIDs := context.Get(r, "ids").([]int)
	// get recipe by ids
	if len(recipeIDs) > 0 {
		recipes, err := getRecipeByIDs(recipeIDs)
		if err != nil {
			return err
		}
		respond.Format(w, r, http.StatusOK, recipes)
		return nil
	}
	// get all recipes

	log.Info(recipeIDs)
	recipe, err := getRecipeByID(1)
	if err != nil {
		return err
	}

	respond.Format(w, r, http.StatusOK, recipe)
	return nil
}

func getRecipeByIDs(recipeIDs []int) ([]*recipe.Recipe, *errors.AppError) {
	var recipes []*recipe.Recipe
	for _, recipeID := range recipeIDs {
		recipe, err := getRecipeByID(recipeID)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func getRecipeByID(recipeID int) (*recipe.Recipe, *errors.AppError) {
	var response recipe.Recipe

	client := &http.Client{
		Timeout: 2 * time.Second,
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

func getID(r *http.Request, keyID string) int {
	key := mux.Vars(r)[keyID]
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Error(err)
	}
	return id
}
