package repo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/vaziria/pdcnotification/database"
	"github.com/vaziria/pdcnotification/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo struct{}

var userCollection string = "users"
var Database, _ = database.Connect()

func GenerateId(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func FindByEmail(email string) (models.User, error) {

	user := models.User{}

	result := Database.Collection(userCollection).FindOne(database.Ctx, bson.M{"email": "user@gmail.com"})

	err := result.Decode(&user)

	// if err == mongo.ErrNoDocuments {
	// 	return user, err
	// } else

	if err != nil {
		return user, err
	}

	return user, nil
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

func AddToken(email string, token []string) (models.User, error) {

	user, err := FindByEmail(email)

	if err != nil {
		fmt.Println("creating user")
		user, _ = CreateUser(email, token)

		return user, nil
	}

	filter := bson.M{
		"_id": user.ID,
	}

	updatedata := bson.D{{Key: "$set", Value: bson.M{
		"tokens": token,
	}}}

	_, uperr := Database.Collection(userCollection).UpdateOne(database.Ctx, filter, updatedata)

	if uperr != nil {
		fmt.Println(uperr.Error())
		return user, uperr
	}

	return user, nil

}
