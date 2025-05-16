package repositories

import (
    "context"
    //"go.mongodb.org/mongo-driver/bson"
    //"go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
    //"log"
    //"os"  

)


type BookRepository struct {
    //collection *mongo.Collection
}

// repositories/book_repository.go
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
    var books []models.Book

    books := []models.Book{
        {
            ID:     "1",
            Title:  "Libro de ejemplo",
            Author: "Autora X",
        },
    }
    // cursor, err := r.collection.Find(context.TODO(), bson.M{})
    // if err != nil {
    //     return nil, err
    // }
    // defer cursor.Close(context.TODO())

    // for cursor.Next(context.TODO()) {
    //     var book models.Book
    //     if err := cursor.Decode(&book); err != nil {
    //         return nil, err
    //     }
    //     books = append(books, book)
    // }

    return books, nil
}