package apigame

import (
	db "admin/database"
	jwt "admin/jwt"
	mapi "admin/models/api"
	mreq "admin/models/api/apirequest"
	mres "admin/models/api/apiresponse"
	muser "admin/models/user"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AuthUser : authentication
func AuthUser(response http.ResponseWriter, request *http.Request) {
	var bodyRequest mreq.AuthRequest
	var con = db.GetConnection()
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&bodyRequest); err != nil {
		log.Println(err.Error())
		jsonString, _ := json.Marshal(mapi.InternalServerError())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)

		return
	}

	user := muser.FindUser(con, bson.M{"Username": bodyRequest.Username})

	if user == (muser.User{}) {
		log.Println("User not found")
		jsonString, _ := json.Marshal(mapi.DataNotFound())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		response.Write(jsonString)
		return
	}
	bodyResponse := mres.AuthResponse{
		Username: bodyRequest.Username,
		DeviceID: bodyRequest.DeviceID,
		Token:    jwt.CreateJWT(bodyRequest.Username, bodyRequest.DeviceID),
		Response: mapi.SuccessMessage(),
	}
	jsonString, _ := json.Marshal(bodyResponse)
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonString)
	return
}

//AddUser : add new user
func AddUser(response http.ResponseWriter, request *http.Request) {
	var bodyRequest muser.User
	var con = db.GetConnection()
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&bodyRequest); err != nil {
		log.Println(err.Error())
		jsonString, _ := json.Marshal(mapi.InternalServerError())
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(jsonString)
		return
	}
	bodyRequest.Created = time.Now()
	bodyRequest.Updated = time.Now()

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

//FindUser : find data user
func FindUser(response http.ResponseWriter, request *http.Request) {
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
