package models

type User struct {
	ID     string   `bson:"_id"`
	Email  string   `bson:"email"`
	Tokens []string `bson:"tokens"`
}
