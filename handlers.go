package main

import (
	"net/http"
)

type FormData struct {
	*Rspv
	Errors []string
}

func WelcomeHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["welcome"].Execute(writer, FormData{
			Rspv:   &Rspv{},
			Errors: []string{},
		})
	}
}

func ListHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["list"].Execute(writer, responses)
	}
}

func FormHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, FormData{
			Rspv:   &Rspv{},
			Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rspv{
			Name:  request.Form.Get("name"),
			Email: request.Form.Get("email"),
			Phone: request.Form.Get("phone"),
		}

		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Укажите своё имя")
		}
		if responseData.Email == "" {
			errors = append(errors, "Укажите свою почту")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Укажите свой номер телефона")
		}

		if len(errors) > 0 {
			templates["form"].Execute(writer, FormData{
				&responseData, errors,
			})
		} else {
			responses = append(responses, &responseData)
			templates["thanks"].Execute(writer, responseData.Name)
		}
	}
}
