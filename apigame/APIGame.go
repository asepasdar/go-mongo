package apigame

import (
	db "admin/database"
	mapi "admin/models/api"
	mgame "admin/models/game"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Progress : add new progress user
func Progress(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		requestObject := struct {
			Score int
			Game  string
			User  string
		}{}
		con := db.GetConnection()

		if err := decoder.Decode(&requestObject); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		objGame, _ := primitive.ObjectIDFromHex(requestObject.Game)
		objUser, _ := primitive.ObjectIDFromHex(requestObject.User)
		var data = mgame.Progress{
			ProgressID: primitive.NewObjectID(),
			GameID:     objGame,
			UserID:     objUser,
			Score:      requestObject.Score,
			Created:    time.Now(),
		}
		if er := data.InsertData(con); er != nil {
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

//Add : add new game
func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		var bodyRequest mgame.Game
		var con = db.GetConnection()
		decoder := json.NewDecoder(request.Body)

		if err := decoder.Decode(&bodyRequest); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

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

//Find : find data game
func Find(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var collectionID = request.URL.Query().Get("id")
		if collectionID != "" {
			var con = db.GetConnection()
			objID, _ := primitive.ObjectIDFromHex(collectionID)
			var data = mgame.FindData(con, bson.M{"_id": objID})

			if data == (mgame.Game{}) {
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
		jsonString, _ := json.Marshal(mapi.DataNotFound())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(jsonString)
		return
	}
}
