package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
    // ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    ID     string `json:"id"` // Temporalmente usamos string
    Title  string `json:"title"`
    ISBN   string `json:"isbn"`
    Author string `json:"author"`
}
