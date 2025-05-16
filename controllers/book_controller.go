package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
    Service services.BookServiceInterface
}

func NewBookController(service services.BookServiceInterface) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
    books, err := c.Service.GetBooks()
    if err != nil {
        http.Error(w, "Error al obtener libros", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Libros obtenidos correctamente",
        "books":   books,
    })
}