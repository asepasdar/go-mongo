package apigame

import (
	db "admin/database"
	mapi "admin/models/api"
	mgame "admin/models/game"
	"encoding/json"
	"net/http"
	"time"
)

//AddUser : add new user
func AddUser(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		var bodyRequest mgame.User
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
