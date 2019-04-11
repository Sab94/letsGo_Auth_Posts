package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Email   string `json:"email"`
	Password        string `json:"password"`
}

type LoginRequest struct {
	Email   string `json:"email"`
	Password        string `json:"password"`
}