package api

import (
	"manigandand-golang-test/pkg/errors"
	"manigandand-golang-test/pkg/recipe"
	"manigandand-golang-test/pkg/respond"
	"net/http"
)

// InitRecipe initiates the recipe enpoints
func InitRecipe() {
	// Get the recipe
	BaseRoutes.Recipes.Handle(
		"", apiHandler(getRecipeHandler),
	).Methods(http.MethodGet)
}

func getRecipeHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	res := new(recipe.Recipe)

	respond.Format(w, r, http.StatusOK, res)
	return nil
}
