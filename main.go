package main

import (
	"net/http"

	"github.com/TodoFromMemory/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("/todo", routers.TodoRoutes())
	http.ListenAndServe(":8000", router)
}
