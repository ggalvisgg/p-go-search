package repositories

import (
    // "context"
    // "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
    // "log"
    // "os"
)

type BookRepository struct {
    // collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
    return &BookRepository{}
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
    books := []models.Book{
        {
            ID:     "1",
            Title:  "Libro de ejemplo",
            Author: "Autora X",
            ISBN:   "1234567890",
        },
    }

    return books, nil
}