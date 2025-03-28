package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// app.notFound(w)
		return
	}
	date := time.Now()

	td := &templateData{
		CurrentYear: date.Year(),
	}

	app.render(w, r, "home.page.tmpl", td)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		// app.notFound(w)
		return
	}
	fmt.Println("login")

	// app.render(w, r, "login.page.tmpl", &templateData{})
}

func (app *application) submitActa(w http.ResponseWriter, r *http.Request) {
	data := r.PostForm

	fmt.Println("data", data)
	if r.URL.Path != "/submit-acta" {
		// app.notFound(w)
		return
	}

	fmt.Println("submit acta")
}
