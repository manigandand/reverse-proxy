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
	"sort"
	"strconv"
	"sync"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// RecipeChanRes holds the response of recipe request
type RecipeChanRes struct {
	Recipe *recipe.Recipe
	Err    *errors.AppError
}

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
	recipe, err := getRecipeFromServer(recipeID)
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
	recipe, err := getRecipeFromServer(1)
	if err != nil {
		return err
	}

	respond.Format(w, r, http.StatusOK, recipe)
	return nil
}

func getRecipeByIDs(recipeIDs []int) ([]*recipe.Recipe, *errors.AppError) {
	var (
		recipes recipe.RecipesSort
		wg      sync.WaitGroup
	)
	totalRecipes := len(recipeIDs)
	done := make(chan bool)
	resultChan := make(chan *RecipeChanRes)
	wg.Add(totalRecipes)
	for _, recipeID := range recipeIDs {
		go getRecipeByID(recipeID, &wg, done, resultChan)
	}
	// read all the response from goroutines
	go func(total int) {
		for i := 0; i < total; i++ {
			select {
			case isDone := <-done:
				if isDone {
					log.Infoln("received done.")
				}
			case res := <-resultChan:
				if res.Err == nil {
					recipes = append(recipes, res.Recipe)
				}
			}
		}
	}(totalRecipes * 2)

	wg.Wait()
	close(done)
	close(resultChan)
	sort.Sort(recipes)

	return recipes, nil
}

func getRecipeByID(recipeID int, wg *sync.WaitGroup, done chan bool, resultChan chan *RecipeChanRes) {
	defer func() {
		done <- true
		wg.Done()
	}()

	recipe, err := getRecipeFromServer(recipeID)
	resultChan <- &RecipeChanRes{
		Recipe: recipe,
		Err:    err,
	}
	return
}

func getRecipeFromServer(recipeID int) (*recipe.Recipe, *errors.AppError) {
	var response recipe.Recipe

	client := &http.Client{
		Timeout: config.ClientTimeout,
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
