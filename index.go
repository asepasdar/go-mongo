package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	api "admin/apigame"
	conf "admin/config"
	c "admin/controllers"
	db "admin/database"
	j "admin/jwt"
	mapi "admin/models/api"

	"github.com/dgrijalva/jwt-go"
)

//File config will be executed first (func init())
func main() {
	db.Initialize(conf.Configuration.Database.URI, conf.Configuration.Database.Name)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", c.Login)
	http.HandleFunc("/dashboard", c.Home)
	http.HandleFunc("/ajaxgamegraph", c.AjaxGraph)

	/* START API REQUEST */
	http.Handle("/auth", middlewarePost(http.HandlerFunc(api.AuthUser)))

	http.Handle("/game", middlewareJWT(middlewareGet(http.HandlerFunc(api.FindGame))))
	http.Handle("/game/add", middlewareJWT(middlewarePost(http.HandlerFunc(api.AddGame))))
	http.Handle("/game/progress", middlewareJWT(middlewarePost(http.HandlerFunc(api.Progress))))

	http.Handle("/user", middlewareJWT(middlewareGet(http.HandlerFunc(api.FindUser))))
	http.Handle("/user/add", middlewareJWT(middlewarePost(http.HandlerFunc(api.AddUser))))
	/* END API REQUEST */

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

func middlewareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		authorizationHeader := request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			jsonString, _ := json.Marshal(mapi.Unauthorized())
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusUnauthorized)
			response.Write(jsonString)
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := j.DecodeJWT(tokenString)

		if err != nil {
			jsonString, _ := json.Marshal(mapi.Unauthorized)
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusUnauthorized)
			response.Write(jsonString)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			jsonString, _ := json.Marshal(mapi.Unauthorized)
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusUnauthorized)
			response.Write(jsonString)
			return
		}

		log.Println(claims)
		log.Println(claims["Username"])

		ctx := context.WithValue(context.Background(), j.KeyInfo, claims)
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)
	})
}
