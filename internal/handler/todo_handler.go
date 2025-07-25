package handler

import (
	"encoding/json"
	"net/http"
	"todo-app/internal/model"
	"todo-app/internal/service"

	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	Service *service.TodoService
}

func(h *TodoHandler)RegisterRoutes(r chi.Router){
	r.Get("/",h.GetAll)
	r.Post("/",h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

func (h *TodoHandler)GetAll(w http.ResponseWriter,r *http.Request){
	todos,err:=h.Service.GetAll()
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}


func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler)GetByID(w http.ResponseWriter,r *http.Request){
	id := chi.URLParam(r, "id")
	
	todo, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler)Update(w http.ResponseWriter,r *http.Request){
	id := chi.URLParam(r, "id")

	var updated model.Todo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated.ID = id

	if err := h.Service.Update(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updated)
}

func (h *TodoHandler)Delete(w http.ResponseWriter,r *http.Request){
	id := chi.URLParam(r, "id")

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}