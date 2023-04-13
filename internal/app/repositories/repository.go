package repositories

import (
	"log"

	"github.com/lastreq/bcraft-test-task/internal/app/models"
)

func (repo *PgDatabase) GetRecipes() (recipes []models.Recipe, err error) {
	sqlStatement := `
SELECT * FROM recipes`

	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return []models.Recipe{}, err
	}

	for rows.Next() {
		var recipe models.Recipe

		if err = rows.Scan(&recipe.RecipeID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.CookingSteps); err != nil {
			return []models.Recipe{}, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (repo *PgDatabase) GetRecipe(recipeID int) (recipe models.Recipe, err error) {
	sqlStatement := `
SELECT * FROM recipes WHERE recipe_id = $1`

	rows, err := repo.DB.Query(sqlStatement, recipeID)
	if err != nil {
		return models.Recipe{}, err
	}

	rows.Next()

	if err = rows.Scan(&recipe.RecipeID, &recipe.Name, &recipe.Description, &recipe.Ingredients, &recipe.CookingSteps); err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (repo *PgDatabase) CreateRecipe(recipe models.Recipe) (err error) {
	sqlStatement := `
INSERT INTO "recipes" ("name", "description", "ingredients", "cooking_steps")
VALUES ($1, $2, $3, $4)`

	_, err = repo.DB.Exec(sqlStatement, recipe.Name, recipe.Description, recipe.Ingredients, recipe.CookingSteps)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PgDatabase) UpdateRecipe(recipe models.Recipe) (err error) {
	sqlStatement := `
UPDATE "recipes" SET name = $2, description = $3, ingredients = $4, cooking_steps = $5 WHERE recipe_id = $1`

	_, err = repo.DB.Exec(sqlStatement, recipe.RecipeID, recipe.Name, recipe.Description, recipe.Ingredients, recipe.CookingSteps)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PgDatabase) DeleteRecipe(recipeID int) (err error) {
	sqlStatement := `
DELETE  FROM recipes WHERE recipe_id = $1`

	_, err = repo.DB.Exec(sqlStatement, recipeID)
	if err != nil {
		log.Println(err.Error())
	}

	return err
}
