package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID        primitive.ObjectID
	Name      string
	Tags      []string
	Cover     string
	Origin    string
	Timestamp int64
	Likes     int64
	Unlikes   int64
	Pics      []string
}
