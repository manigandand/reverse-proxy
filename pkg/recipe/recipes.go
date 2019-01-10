package recipe

// Recipe holds the information about the recipe properties
type Recipe struct {
	ID          int          `json:"id,string"`
	Name        string       `json:"name"`
	Headline    string       `json:"headline"`
	Description string       `json:"description"`
	Difficulty  int          `json:"difficulty"`
	PrepTime    string       `json:"prepTime"`
	ImageLink   string       `json:"imageLink"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient holds the ingredient properties
type Ingredient struct {
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}

// RecipesSort should sort given recipe container values by preparation time
// of repos in ascending order.
type RecipesSort []*Recipe

func (r RecipesSort) Len() int           { return len(r) }
func (r RecipesSort) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RecipesSort) Less(i, j int) bool { return r[i].PrepTime < r[j].PrepTime }
