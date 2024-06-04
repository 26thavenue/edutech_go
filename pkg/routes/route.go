package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/26thavenue/edutech_go/pkg/controllers"
)

func BookRoutes() chi.Router {
    r := chi.NewRouter()
    r.Get("/", controllers.ListBooks)
    r.Post("/", controllers.CreateBook)
    r.Get("/{id}", controllers.GetBooks)
    r.Put("/{id}", controllers.UpdateBook)
    r.Delete("/{id}", controllers.DeleteBook)
    return r
}