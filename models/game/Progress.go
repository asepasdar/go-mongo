package mgame

import (
	"context"
	"time"

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

//InsertData : insert new admin data
func (p Progress) InsertData(con *mongo.Database) error {
	var ctx = context.Background()
	var err error
	_, err = con.Collection("Progress").InsertOne(ctx, p)

	return err
}
