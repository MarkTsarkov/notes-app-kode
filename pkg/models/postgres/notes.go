package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/marktsarkov/notes-app-kode/pkg/models"
)

type NoteModel struct {
	DB *pgx.Conn
}

func (m *NoteModel) Insert(username, note string) error {
	req := `INSERT INTO users (user_name, note) VALUES ($1, $2)`

	_, err := m.DB.Exec(context.Background(), req, username, note)
	if err!=nil{
		return err
	}

	return nil
}

func (m *NoteModel) Get(user string) ([]*models.Note, error) {
	req := `SELECT note FROM users WHERE user_name = $1`
	
	rows, err := m.DB.Query(context.Background(), req, user)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()
	
	var notes []*models.Note

	for rows.Next() {
		n := &models.Note{}

		err = rows.Scan(&n.Note)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Println(notes) 
	return notes, nil
}


 
 
	