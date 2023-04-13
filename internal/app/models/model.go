package models

type Recipe struct {
	RecipeID     int64  `json:"recipe_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Ingredients  string `json:"ingredients,omitempty"`
	CookingSteps string `json:"cooking_steps,omitempty"`
}
