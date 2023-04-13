package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router     *mux.Router
	controller controller
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}

type controller interface {
	GetRecipes(w http.ResponseWriter, r *http.Request)
	GetRecipe(w http.ResponseWriter, r *http.Request)
	CreateRecipe(w http.ResponseWriter, r *http.Request)
	UpdateRecipe(w http.ResponseWriter, r *http.Request)
	DeleteRecipe(w http.ResponseWriter, r *http.Request)
}

func (r *Router) GetRouter() http.Handler {
	return r.router
}

func New() *Router {
	router := mux.NewRouter()

	r := Router{
		router: router,
	}

	return &r
}

func (r *Router) AppendRoutes(ctr controller) {
	routes := []Route{
		{
			Name:    "recipes",
			Path:    "/recipes",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(ctr.GetRecipes),
		},
		{
			Name:    "recipe",
			Path:    "/recipe",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(ctr.GetRecipe),
		},
		{
			Name:    "recipe",
			Path:    "/recipe",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(ctr.CreateRecipe),
		},
		{
			Name:    "recipe",
			Path:    "/recipe",
			Method:  http.MethodPut,
			Handler: http.HandlerFunc(ctr.UpdateRecipe),
		},
		{
			Name:    "recipe",
			Path:    "/recipe",
			Method:  http.MethodDelete,
			Handler: http.HandlerFunc(ctr.DeleteRecipe),
		},
	}

	for _, route := range routes {
		r.router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
