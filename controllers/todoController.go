package controllers

import (
	"fmt"
	"net/http"
)

type TodoController struct{}

func (t TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	username := r.Context().Value("username").(string)
	w.Write([]byte(fmt.Sprintf("User id is %v and username is %v", user, username)))
}
