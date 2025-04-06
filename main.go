package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rspv struct {
	Name, Email, Phone string
}

var responses = make([]*Rspv, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	LoadTemplates()

	http.HandleFunc("/", WelcomeHandler)
	http.HandleFunc("/list", ListHandler)
	http.HandleFunc("/form", FormHandler)

	fmt.Println("Сервер запущен на http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
