package controlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lastreq/bcraft-test-task/internal/app/models"
)

func (ctr *Controller) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := ctr.svc.GetRecipes()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	recipesJSONBytes, err := json.Marshal(recipes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(recipesJSONBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	return
}

func (ctr *Controller) GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	recipe, err := ctr.svc.GetRecipe(recipeID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	recipeJSONBytes, err := json.Marshal(recipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(recipeJSONBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	return
}

func (ctr *Controller) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	var recipe models.Recipe

	err = json.Unmarshal(bodyBytes, &recipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)

		return
	}

	err = ctr.svc.CreateRecipe(recipe)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)

	return
}

func (ctr *Controller) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	var recipe models.Recipe

	err = json.Unmarshal(bodyBytes, &recipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)

		return
	}

	err = ctr.svc.UpdateRecipe(recipe)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	return
}

func (ctr *Controller) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = ctr.svc.DeleteRecipe(recipeID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	return
}
