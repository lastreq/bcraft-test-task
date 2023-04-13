package main

import (
	"log"
	"net/http"

	"github.com/lastreq/bcraft-test-task/internal/app/controlers"
	"github.com/lastreq/bcraft-test-task/internal/app/repositories"
	"github.com/lastreq/bcraft-test-task/internal/app/router"
	"github.com/lastreq/bcraft-test-task/internal/app/services"
)

func main() {
	repo := repositories.New()

	svc := services.New(repo)

	ctr := controlers.New(svc)

	r := router.New()

	r.AppendRoutes(ctr)

	err := http.ListenAndServe(":8000", r.GetRouter())
	if err != nil {
		log.Println(err)
	}
}
