package main

import (
	"fmt"
	"html/template"
)

func LoadTemplates() {
	templateNames := [4]string{
		"welcome", "form", "thanks", "list",
	}
	for index, name := range templateNames {
		t, err := template.ParseFiles("templates/layout.html", "templates/"+name+".html")
		if err != nil {
			panic(err)
		}
		templates[name] = t
		fmt.Println("Loaded template", index, name)
	}
}
