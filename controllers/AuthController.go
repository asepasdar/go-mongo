package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	db "admin/database"
	madmin "admin/models/admin"
)

type templateData map[string]interface{}

//Login : handle request login page
func Login(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getLogin(response, request)

	case http.MethodPost:
		postLogin(response, request)

	default:
		http.Error(response, "", http.StatusBadRequest)
	}
}

func getLogin(response http.ResponseWriter, request *http.Request) {
	var data = templateData{
		"title": "Login Page | Golang & MongoDB",
	}
	var tmpl = template.Must(template.ParseFiles(
		"views/shared/login/_HeaderLogin.html",
		"views/shared/login/_FooterLogin.html",
		"views/auth/Login.html",
	))
	err := tmpl.ExecuteTemplate(response, "login", data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func postLogin(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	var con = db.GetConnection()
	var criteria = map[string]interface{}{
		"Username": request.PostFormValue("Username"),
		"Password": request.PostFormValue("Password"),
	}
	var findData = madmin.FindData(con, criteria)

	fmt.Println(criteria, " ", findData)
	if findData != (madmin.Admin{}) {
		fmt.Println(findData)
		http.Redirect(response, request, "/dashboard", http.StatusFound)
	}
}
