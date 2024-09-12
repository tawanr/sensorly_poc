package main

import "net/http"

func (app *application) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	app.notFoundResponse(w, r)
}
