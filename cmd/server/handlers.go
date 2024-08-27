package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/marktsarkov/notes-app-kode/pkg/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

    w.Write([]byte("aloha"))
}

func (app *application) authorization (user, pass string, w http.ResponseWriter) bool {
	access := auth(user, pass)
	if !access {
		app.clientError(w, http.StatusForbidden)
		return false
	}
	return true
}

func (app *application) ShowNote(w http.ResponseWriter, r *http.Request) {
	var notes models.Note
	
	decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&notes)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

	if !app.authorization(notes.User, notes.Password, w){
		return
	}

    s, err := app.notes.Get(notes.User)
    if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	
	jsonResponse, err := json.Marshal(s)
    if err != nil {
        app.serverError(w, err)
        return
    }
    w.Write(jsonResponse)
}

func (app *application) CreateNote(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

	var notes models.Note
	
	decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&notes)
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

	if !app.authorization(notes.User, notes.Password, w){
		return
	}

    spellCheckResponses, err := checkSpelling(notes.Note)
    if err != nil {
        app.serverError(w, err)
        return
    }
    if len(spellCheckResponses) > 0 {
		jsonResponse, err := json.Marshal(spellCheckResponses)
		if err != nil {
			app.serverError(w, err)
			return
		}
		w.Write(jsonResponse)
		return
    }

	err = app.notes.Insert(notes.User, notes.Note)
	if err != nil {
		app.serverError(w, err)
		return
	}


    w.Write([]byte("Заметка создана"))
}