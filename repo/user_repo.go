package repo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/vaziria/pdcnotification/database"
	"github.com/vaziria/pdcnotification/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct{}

var userCollection string = "users"
var Database, _ = database.Connect()

func GenerateId(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func FindByEmail(email string) (models.User, bool) {

	user := models.User{}

	err := Database.Collection(userCollection).FindOne(database.Ctx, bson.M{email: email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		log.Fatal(err)
	} else if err != nil {
		log.Fatal(err)
	}

	return user, false
}

func CreateUser(email string, tokens []string) (models.User, bool) {

	// inserting data
	oid := GenerateId(email)
	user := models.User{
		ID:     oid,
		Email:  email,
		Tokens: tokens,
	}

	_, err := Database.Collection(userCollection).InsertOne(database.Ctx, user)
	if err != nil {
		return user, true
	}

	return user, false
}

func AddToken(email string, token []string) (models.User, bool) {

	user, err := FindByEmail(email)

	if err {
		fmt.Println("creating user")
		user, _ = CreateUser(email, token)

		return user, false
	}

	filter := bson.M{
		"_id": user.ID,
	}

	updatedata := bson.D{{Key: "$set", Value: models.User{
		Tokens: token,
	}}}

	_, uperr := Database.Collection(userCollection).UpdateOne(database.Ctx, filter, updatedata)

	if uperr != nil {
		log.Fatal("Error on updating token", err)
	}

	return user, false

}
