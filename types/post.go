package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content string `json:"content"`
	Poster   primitive.ObjectID `json:"poster"`
}