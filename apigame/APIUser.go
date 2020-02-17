package apigame

import (
	db "admin/database"
	mapi "admin/models/api"
	muser "admin/models/user"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddUser : add new user
func AddUser(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		var bodyRequest muser.User
		var con = db.GetConnection()
		decoder := json.NewDecoder(request.Body)

		if err := decoder.Decode(&bodyRequest); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		bodyRequest.Created = time.Now()
		bodyRequest.Updated = time.Now()

		if er := bodyRequest.InsertData(con); er != nil {
			jsonString, _ := json.Marshal(mapi.MongoExecutionFailed())
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusCreated)
			response.Write(jsonString)
			return
		}

		jsonString, _ := json.Marshal(mapi.SuccessMessage())
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonString)
	}
}

//FindUser : find data user
func FindUser(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var con = db.GetConnection()
		var collectionID = request.URL.Query().Get("id")
		if collectionID != "" {

			objID, _ := primitive.ObjectIDFromHex(collectionID)
			var data = muser.FindUser(con, bson.M{"_id": objID})

			if data == (muser.User{}) {
				jsonString, _ := json.Marshal(mapi.DataNotFound())
				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write(jsonString)
				return
			}

			jsonString, _ := json.Marshal(data)
			response.Header().Set("Content-Type", "application/json")
			response.Write(jsonString)
			return
		}

		var criteria = map[string]interface{}{}
		var data = muser.FindAllUser(con, criteria)
		jsonString, _ := json.Marshal(data)
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonString)
		return
	}
}
