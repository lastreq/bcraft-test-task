package controlers

import (
	"github.com/lastreq/bcraft-test-task/internal/app/models"
)

type Controller struct {
	svc service
}
type service interface {
	GetRecipes() (recipes []models.Recipe, err error)
	GetRecipe(recipeID int) (recipe models.Recipe, err error)
	CreateRecipe(recipe models.Recipe) (err error)
	UpdateRecipe(recipe models.Recipe) (err error)
	DeleteRecipe(recipeID int) (err error)
}

func New(svc service) *Controller {
	ctr := &Controller{svc: svc}

	return ctr
}
