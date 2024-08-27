package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Note struct {
	ID      int
	User    string
	Note	string
}