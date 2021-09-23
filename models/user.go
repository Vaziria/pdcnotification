package models

type User struct {
	Email string   `bson:"name"`
	Token []string `bson:"token"`
}
