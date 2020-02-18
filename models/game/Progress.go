package mgame

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Progress : progress object
type Progress struct {
	ProgressID primitive.ObjectID `bson:"_id, omitempty"`
	GameID     primitive.ObjectID `bson:"GameID, omitempty"`
	UserID     primitive.ObjectID `bson:"UserID, omitempty"`
	Score      int                `bson:"Score"`
	Created    time.Time          `bson:"Created"`
}

//Leaderboard : leaderboard object
type Leaderboard struct {
	ProgressID primitive.ObjectID `bson:"_id, omitempty"`
	GameID     primitive.ObjectID `bson:"GameID, omitempty"`
	UserID     primitive.ObjectID `bson:"UserID, omitempty"`
	Score      int                `bson:"Score"`
	Created    time.Time          `bson:"Created"`
	Game       []struct {
		GameID      primitive.ObjectID `bson:"_id, omitempty"`
		Title       string             `bson:"Title"`
		Description string             `bson:"Description"`
	}
	User []struct {
		UserID   primitive.ObjectID `bson:"_id, omitempty"`
		Username string             `bson:"Username"`
		DeviceID string             `bson:"DeviceID"`
		Created  time.Time          `bson:"Created"`
		Updated  time.Time          `bson:"Updated"`
	}
}

//InsertData : insert new admin data
func (p Progress) InsertData(con *mongo.Database) error {
	var ctx = context.Background()
	var err error
	_, err = con.Collection("Progress").InsertOne(ctx, p)
	return err
}

//DateString : Return date time to string
func (l Leaderboard) DateString() string {
	var result = l.Created.Format("02 January 2006 - 15:04")
	return result
}

//FindAllProgress : select all data from progress depends on criteria
func FindAllProgress(con *mongo.Database, criteria map[string]interface{}) []Progress {
	var ctx = context.Background()

	var qresult, err = con.Collection("Progress").Find(ctx, criteria)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer qresult.Close(ctx)

	result := make([]Progress, 0)
	for qresult.TryNext(ctx) {
		var row Progress
		err := qresult.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return result
}

//FindLeaderboard : find leaderboard
func FindLeaderboard(con *mongo.Database, criteria *primitive.ObjectID) []Leaderboard {
	var ctx = context.Background()
	var cobapipeline = []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "Game",
				"localField":   "GameID",
				"foreignField": "_id",
				"as":           "Game",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "User",
				"localField":   "UserID",
				"foreignField": "_id",
				"as":           "User",
			},
		},
		bson.M{
			"$match": bson.M{
				"GameID": criteria,
			},
		},
	}

	csr, err := con.Collection("Progress").Aggregate(ctx, cobapipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Leaderboard, 0)
	for csr.Next(ctx) {
		var row Leaderboard
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	return result
}
