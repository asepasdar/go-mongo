package mgame

import (
	"context"
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
