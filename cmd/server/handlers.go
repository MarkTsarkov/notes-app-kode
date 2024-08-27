package main

import (
	"errors"
	"fmt"
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

func (app *application) ShowNote(w http.ResponseWriter, r *http.Request) {
    user := r.URL.Query().Get("user")

    s, err := app.notes.Get(user)
    if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
 
    for _, note := range s {
        fmt.Fprintf(w, "%v\n", note)
    }
}

func (app *application) CreateNote(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    // Создаем несколько переменных, содержащих тестовые данные. Мы удалим их позже.
	username := "boris"
	note := "Молоко, Каша, Рис"
 
	// Передаем данные в метод SnippetModel.Insert(), получая обратно
	// ID только что созданной записи в базу данных.
	err := app.notes.Insert(username, note)
	if err != nil {
		app.serverError(w, err)
		return
	}


    w.Write([]byte("Создание новой заметки..."))
}