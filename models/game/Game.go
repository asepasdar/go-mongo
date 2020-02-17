package mgame

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//Game : object game
type Game struct {
	Title       string `bson:"Title"`
	Description string `bson:"Description"`
}

//InsertData : insert new game data
func (g Game) InsertData(con *mongo.Database) error {
	var ctx = context.Background()
	var err error
	_, err = con.Collection("Game").InsertOne(ctx, g)

	return err
}

//FindData : find admin data
func FindData(con *mongo.Database, criteria map[string]interface{}) Game {
	var ctx = context.Background()
	var data Game

	var myResult = con.Collection("Game").FindOne(ctx, criteria)
	myResult.Decode(&data)

	return data
}
