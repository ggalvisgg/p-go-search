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
    collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("library").Collection("books")
    return &BookRepository{collection}
}



// repositories/book_repository.go
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
    var books []models.Book
    cursor, err := r.collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var book models.Book
        if err := cursor.Decode(&book); err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    return books, nil
}