package mgame

import (
	"context"
	"fmt"
	"log"

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

//FindGame : find game data
func FindGame(con *mongo.Database, criteria map[string]interface{}) Game {
	var ctx = context.Background()
	var data Game

	var myResult = con.Collection("Game").FindOne(ctx, criteria)
	myResult.Decode(&data)

	return data
}

//FindAllGame : select all data from game depends on criteria
func FindAllGame(con *mongo.Database, criteria map[string]interface{}) []Game {
	var ctx = context.Background()

	var qresult, err = con.Collection("Game").Find(ctx, criteria)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := make([]Game, 0)
	for qresult.TryNext(ctx) {
		var row Game
		err := qresult.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return result
}
