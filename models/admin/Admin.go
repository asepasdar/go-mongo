package madmin

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//Admin : object user admin
type Admin struct {
	Username string `bson:"Username"`
	Password string `bson:"Password"`
}

//InsertData : insert new admin data
func InsertData(con *mongo.Database, data []interface{}) error {
	var ctx = context.Background()
	var err error
	_, err = con.Collection("Admin").InsertMany(ctx, data)

	return err
}

//FindData : find admin data
func FindData(con *mongo.Database, criteria map[string]interface{}) Admin {
	var ctx = context.Background()
	var data Admin

	var myResult = con.Collection("Admin").FindOne(ctx, criteria)
	myResult.Decode(&data)

	return data
}
