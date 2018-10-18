package api

import (
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net/http"

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
	res := new(recipe.Recipe)
	ids := context.Get(r, "ids").([]int)

	log.Info(ids)
	respond.Format(w, r, http.StatusOK, res)
	return nil
}
