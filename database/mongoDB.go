package database

import (
	madmin "admin/models/admin"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var connection *mongo.Database

/* PUBLIC FUNCTION RIGHT HERE */

//Initialize : create connection to MongoDB (localhost)
func Initialize(dbURI string, dbName string) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println("try connect to database :  success")
	connection = client.Database(dbName)
	createAdmin()
}

//GetConnection :  get connection
func GetConnection() *mongo.Database {
	return connection
}

/* END PUBLIC FUNCTION */

/* PRIVATE FUNCTION */

func createAdmin() {
	var search = map[string]interface{}{
		"Username": "admin",
	}
	var result = madmin.FindData(connection, search)
	if result == (madmin.Admin{}) {
		var data = []interface{}{
			madmin.Admin{
				Username: "admin",
				Password: "admin",
			},
		}
		madmin.InsertData(connection, data)
	}

}

/* END PRIVATE FUNCTION */
