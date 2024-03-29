package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Category  string             `bson:"category,omitempty"`
	UserId    string             `bson:"userid,omitempty"`
	Published string             `bson:"published,omitempty"`
	Updated   string             `bson:"updated,omitempty"`
}
