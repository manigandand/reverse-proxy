package schemas

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
