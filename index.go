package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	api "admin/apigame"
	conf "admin/config"
	c "admin/controllers"
	db "admin/database"
	mapi "admin/models/api"
)

//File config will be executed first (func init())
func main() {
	db.Initialize(conf.Configuration.Database.URI, conf.Configuration.Database.Name)

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

	server := new(http.Server)
	server.ReadTimeout = conf.Configuration.Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration.Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Configuration.Server.Port)

	log.Println("server started at", conf.Configuration.Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

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
