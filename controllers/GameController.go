package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	mgame "admin/models/game"
)

//Home : handle request to dashboard page
func Home(response http.ResponseWriter, request *http.Request) {
	var data = templateData{
		"title": "Dashboard Page | Golang & MongoDB",
	}
	var tmpl = template.Must(template.ParseFiles(
		"views/shared/dashboard/_ContentHeader.html",
		"views/shared/dashboard/_ContentTop.html",
		"views/shared/dashboard/_ContentLeft.html",
		"views/shared/dashboard/_ContentFooter.html",
		"views/shared/dashboard/_ScriptDashboard.html",
		"views/shared/dashboard/_EndDashboard.html",
		"views/shared/dashboard/_StartDashboard.html",
		"views/game/Dashboard.html",
	))
	err := tmpl.ExecuteTemplate(response, "dashboard", data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

//AjaxGraph : ajax request return data set for graph on dashboard page
func AjaxGraph(response http.ResponseWriter, request *http.Request) {
	var result []mgame.AjaxGameGraph
	for i := 0; i < 3; i++ {
		var row = mgame.AjaxGameGraph{
			GameID:   i + 1,
			Title:    fmt.Sprint("Game", i),
			Period:   request.FormValue("Period"),
			Data:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			Category: []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		}
		result = append(result, row)
	}

	jsonString, _ := json.Marshal(result)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Write(jsonString)

	return
}
