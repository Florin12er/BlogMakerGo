package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
    ID      primitive.ObjectID `bson:"_id,omitempty"`
    Title   string             `bson:"title"`
    Link    string             `bson:"link"`
    Content string             `bson:"content"`
}

