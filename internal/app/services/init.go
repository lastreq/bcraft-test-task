package services

import (
	"github.com/lastreq/bcraft-test-task/internal/app/models"
)

type Service struct {
	repo repository
}

type repository interface {
	GetRecipes() (recipes []models.Recipe, err error)
	GetRecipe(recipeID int) (recipe models.Recipe, err error)
	CreateRecipe(recipe models.Recipe) (err error)
	UpdateRecipe(recipe models.Recipe) (err error)
	DeleteRecipe(recipeID int) (err error)
}

func New(repo repository) *Service {
	svc := Service{repo: repo}

	return &svc
}
