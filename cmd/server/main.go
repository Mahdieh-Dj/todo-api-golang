package main

import (
	"log"
	"net/http"
	"todo-app/internal/handler"
	"todo-app/internal/model"
	"todo-app/internal/repository"
	"todo-app/internal/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=123456 dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Todo{})

	repo := &repository.GormRepository{DB: db}
	svc := &service.TodoService{Repo: repo}
	todoHandler := &handler.TodoHandler{Service: svc}

	r := chi.NewRouter()
	r.Route("/todos", todoHandler.RegisterRoutes)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}