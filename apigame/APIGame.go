package apigame

import (
	db "admin/database"
	mapi "admin/models/api"
	mgame "admin/models/game"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Progress : add new progress user
func Progress(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	requestObject := struct {
		Score int
		Game  string
		User  string
	}{}
	con := db.GetConnection()

	if err := decoder.Decode(&requestObject); err != nil {
		log.Println(err.Error())
		jsonString, _ := json.Marshal(mapi.ErrorParsingBody())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)
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
		log.Println(er.Error())
		jsonString, _ := json.Marshal(mapi.MongoExecutionFailed())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)
		return
	}

	jsonString, _ := json.Marshal(mapi.SuccessMessage())
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonString)
}

//AddGame : add new game
func AddGame(response http.ResponseWriter, request *http.Request) {
	var bodyRequest mgame.Game
	var con = db.GetConnection()
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&bodyRequest); err != nil {
		log.Println(err.Error())
		jsonString, _ := json.Marshal(mapi.ErrorParsingBody())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)
		return
	}

	if er := bodyRequest.InsertData(con); er != nil {
		log.Println(er.Error())
		jsonString, _ := json.Marshal(mapi.MongoExecutionFailed())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)
		return
	}

	jsonString, _ := json.Marshal(mapi.SuccessMessage())
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonString)
}

//FindGame : find data game
func FindGame(response http.ResponseWriter, request *http.Request) {
	var con = db.GetConnection()
	var collectionID = request.URL.Query().Get("id")
	if collectionID != "" {
		objID, _ := primitive.ObjectIDFromHex(collectionID)
		var data = mgame.FindGame(con, bson.M{"_id": objID})

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

	var criteria = map[string]interface{}{}
	var data = mgame.FindAllGame(con, criteria)
	jsonString, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonString)
	return
}
