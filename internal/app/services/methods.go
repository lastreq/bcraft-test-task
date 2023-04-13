package services

import (
	"github.com/lastreq/bcraft-test-task/internal/app/models"
)

func (svc *Service) GetRecipes() (recipes []models.Recipe, err error) {
	recipes, err = svc.repo.GetRecipes()
	if err != nil {
		return []models.Recipe{}, err
	}

	return recipes, nil
}

func (svc *Service) GetRecipe(recipeID int) (recipe models.Recipe, err error) {
	recipe, err = svc.repo.GetRecipe(recipeID)
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, err
}

func (svc *Service) CreateRecipe(recipe models.Recipe) (err error) {
	err = svc.repo.CreateRecipe(recipe)
	if err != nil {
		return err
	}

	return err
}

func (svc *Service) UpdateRecipe(recipe models.Recipe) (err error) {
	err = svc.repo.UpdateRecipe(recipe)
	if err != nil {
		return err
	}

	return err
}

func (svc *Service) DeleteRecipe(recipeID int) (err error) {
	err = svc.repo.DeleteRecipe(recipeID)
	if err != nil {
		return err
	}

	return err
}
