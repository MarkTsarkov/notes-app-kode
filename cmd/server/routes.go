package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", app.Home)
    mux.HandleFunc("/show", app.ShowNote)
    mux.HandleFunc("/create", app.CreateNote)

	return mux
}