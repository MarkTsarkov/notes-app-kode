package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Note struct {
	ID     	 int 	`json:"id"`
	User   	 string	`json:"user"`
	Password string	`json:"password"`
	Note	 string	`json:"note"`
}