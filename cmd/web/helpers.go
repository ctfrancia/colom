package main

import (
	"bytes"
	"log"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		log.Printf("The template %s does not exist", name)
		// app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		log.Println(err)
		// app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	// Here we are adding the default value to our sessions
	// td.Flash = app.session.PopString(r, "flash")
	// td.CSRFToken = nosurf.Token(r)
	// td.CurrentYear = time.Now().Year()
	// td.IsAuthenticated = app.isAuthenticated(r)
	return td
}
