# 📝 Todo API - Go (Golang) + PostgreSQL

A simple, modular **RESTful Todo API** built in Go using Clean Architecture principles, PostgreSQL for persistence, and Chi router for HTTP routing.

---

## 📦 Features

- Full CRUD operations for todos
- Layered clean architecture (handler → service → repository)
- GORM ORM + PostgreSQL
- Dependency injection and modular design
- Docker Compose for local setup

## 🚀 Getting Started

docker compose up -d
go mod init todo-api-go
go get github.com/go-chi/chi/v5
go get gorm.io/gorm
go get gorm.io/driver/postgres
go run cmd/server/main.go
