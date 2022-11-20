package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/todos", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Get all todos"))
	})
	
	router.Post("/todo/{id}", func(w http.ResponseWriter, req *http.Request) {

		id := chi.URLParam(req, "id")
		w.Write([]byte(fmt.Sprintf("Create a todo with id:  %v", id)))
	})

	http.ListenAndServe(":8000", router)
}

	// router.Mount("/todo", routers.TodoRoutes())