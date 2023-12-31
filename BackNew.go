package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type PageVariables struct {
	Title string
}

type ResponseStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var newUserChannel = make(chan struct{}, 1)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/Script.js", jsFile)
	http.HandleFunc("/loginPage", LoginPage)
	http.HandleFunc("/AdminPage", AdminPage)
	http.HandleFunc("/login", LoginHandler)       //get
	http.HandleFunc("/register", RegisterHandler) //post
	http.HandleFunc("/admin", AdminHandler)       //post
	http.HandleFunc("/admin/all", AdminAll)       //get
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func jsFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("Script.js")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "Admin Page",
	}

	tmpl, err := template.ParseFiles("AdminPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "Login Page",
	}

	tmpl, err := template.ParseFiles("LoginPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "Registration Page",
	}

	tmpl, err := template.ParseFiles("Registration.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sendJSONResponse(w http.ResponseWriter, response ResponseStatus) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	w.Write(responseJSON)
}
