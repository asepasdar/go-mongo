package database

import (
	madmin "admin/models/admin"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var connection *mongo.Database

/* PUBLIC FUNCTION RIGHT HERE */

//Initialize : create connection to MongoDB (localhost)
func Initialize() {
	fmt.Print("try connect to database : ")
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
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
	fmt.Println("success")
	connection = client.Database("web-golang")
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
