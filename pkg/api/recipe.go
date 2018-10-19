package api

import (
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/proxy"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net/http"
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
	page := respond.NewPage(r)
	recipes, isEOF := getAllRecipes(page)

	respond.Paginate(w, r, recipes, page, isEOF, len(recipes))
	return nil
}

func getAllRecipes(page *respond.Page) ([]*recipe.Recipe, bool) {
	var (
		recipes                 recipe.RecipesSort
		wg                      sync.WaitGroup
		isEOF                   bool
		maxConcurrentGoroutines = 5
		bufferChan              = make(chan struct{}, maxConcurrentGoroutines)
		done                    = make(chan bool)
		resultChan              = make(chan *RecipeChanRes)
	)
	for i := 0; i < maxConcurrentGoroutines; i++ {
		bufferChan <- struct{}{}
	}

	// read from chanels
	go func(total int) {
		for i := 0; i < total; i++ {
			select {
			case isDone := <-done:
				if isDone {
					bufferChan <- struct{}{}
				}
			case res := <-resultChan:
				if res.Err == nil {
					recipes = append(recipes, res.Recipe)
				} else {
					if res.Err.IsStatusNotFound() {
						isEOF = true
					}
				}
			}
		}
	}(page.Limit * 2)

	recipeID := page.Offset
	wg.Add(page.Limit)
	for i := 0; i < page.Limit; i++ {
		<-bufferChan
		recipeID++
		go getRecipeByID(recipeID, &wg, done, resultChan)
	}
	wg.Wait()
	close(done)
	close(resultChan)

	return recipes, isEOF
}

func getRecipeByIDs(recipeIDs []int) ([]*recipe.Recipe, *errors.AppError) {
	var (
		recipes    recipe.RecipesSort
		wg         sync.WaitGroup
		done       = make(chan bool)
		resultChan = make(chan *RecipeChanRes)
	)
	totalRecipes := len(recipeIDs)
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
	recipe, err := proxy.GetRecipe(recipeID)
	resultChan <- &RecipeChanRes{
		Recipe: recipe,
		Err:    err,
	}
	return
}

func getID(r *http.Request, keyID string) int {
	key := mux.Vars(r)[keyID]
	id, err := strconv.Atoi(key)
	if err != nil {
		log.Error(err)
	}
	return id
}

// getRecipeHandler
func getRecipeHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	recipeID := getID(r, "recipe-id")
	recipe, err := proxy.GetRecipe(recipeID)
	if err != nil {
		return err
	}
	respond.Format(w, r, http.StatusOK, recipe)
	return nil
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return errors.BadRequest(err.Error())
	// }
	// r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// proxy.ReverseProxy(w, r, recipeID)
}
