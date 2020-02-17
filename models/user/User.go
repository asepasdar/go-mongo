package muser

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//User : object game
type User struct {
	Username string    `bson:"Username"`
	DeviceID string    `bson:"DeviceID"`
	Created  time.Time `bson:"Created"`
	Updated  time.Time `bson:"Updated"`
}

//InsertData : insert new user data
func (g User) InsertData(con *mongo.Database) error {
	var ctx = context.Background()
	var err error
	_, err = con.Collection("User").InsertOne(ctx, g)

	return err
}

//FindUser : find user data
func FindUser(con *mongo.Database, criteria map[string]interface{}) User {
	var ctx = context.Background()
	var data User

	var myResult = con.Collection("User").FindOne(ctx, criteria)
	myResult.Decode(&data)

	return data
}

//FindAllUser : select all data from user depends on criteria
func FindAllUser(con *mongo.Database, criteria map[string]interface{}) []User {
	var ctx = context.Background()

	var qresult, err = con.Collection("User").Find(ctx, criteria)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := make([]User, 0)
	for qresult.TryNext(ctx) {
		var row User
		err := qresult.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return result
}
