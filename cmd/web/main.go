package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	fmt.Println("Hello, World!")
	templateCache, err := newTemplateCache("./ui/html")
	if nil != err {
		log.Fatal(err)
	}
	app := application{templateCache}
	srv := &http.Server{
		Addr:         ":4000",
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	err = srv.ListenAndServe()
	if nil != err {
		log.Fatal(err)
	}
}
