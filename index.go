package main

import (
	"fmt"
	"net/http"

	api "admin/apigame"
	c "admin/controllers"
	db "admin/database"
)

func main() {
	db.Initialize()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", c.Login)
	http.HandleFunc("/dashboard", c.Home)
	http.HandleFunc("/ajaxgamegraph", c.AjaxGraph)

	http.HandleFunc("/game", api.FindGame)
	http.HandleFunc("/game/add", api.Add)
	http.HandleFunc("/game/progress", api.Progress)

	http.HandleFunc("/user", api.FindUser)
	http.HandleFunc("/user/add", api.AddUser)

	fmt.Println("server started at localhost:9000")
	var address = "localhost:9000"
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
