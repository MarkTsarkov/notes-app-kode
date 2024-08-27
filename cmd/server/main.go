package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/marktsarkov/notes-app-kode/pkg/models/postgres"
)

type application struct {
    infoLog     *log.Logger
    errorLog    *log.Logger
    notes       *postgres.NoteModel
}

func main() {
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

    PG_DSN := "host=localhost port=54321 dbname=users user=admin password=admin sslmode=disable"
    conn, err := pgx.Connect(context.Background(), PG_DSN)
    if err != nil {
        errorLog.Fatal(err)
    }

    app := &application{
        infoLog:  infoLog,
		errorLog: errorLog,
        notes:    &postgres.NoteModel{DB: conn},
	}

    srv := &http.Server{
        Addr:     ":4000",
        Handler:  app.routes(),
        ErrorLog: errorLog,
    }

    infoLog.Println("Запуск сервера на http://127.0.0.1:4000") //ADRESS=8080
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}
