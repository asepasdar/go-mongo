package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "admin/apigame"
	c "admin/controllers"
	db "admin/database"
	mapi "admin/models/api"
)

func main() {
	db.Initialize()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", c.Login)
	http.HandleFunc("/dashboard", c.Home)
	http.HandleFunc("/ajaxgamegraph", c.AjaxGraph)

	http.Handle("/auth", middlewarePost(http.HandlerFunc(api.AuthUser)))

	http.Handle("/game", middlewareGet(http.HandlerFunc(api.FindGame)))
	http.Handle("/game/add", middlewarePost(http.HandlerFunc(api.AddGame)))
	http.Handle("/game/progress", middlewarePost(http.HandlerFunc(api.Progress)))

	http.Handle("/user", middlewareGet(http.HandlerFunc(api.FindUser)))
	http.Handle("user/add", middlewarePost(http.HandlerFunc(api.AddUser)))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}

func middlewareGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != "GET" {
			jsonString, _ := json.Marshal(mapi.WrongHTTPMethod())
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusNotAcceptable)
			response.Write(jsonString)
			return
		}
		next.ServeHTTP(response, request)
	})
}

func middlewarePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			jsonString, _ := json.Marshal(mapi.WrongHTTPMethod())
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusNotAcceptable)
			response.Write(jsonString)
			return
		}
		next.ServeHTTP(response, request)
	})
}
