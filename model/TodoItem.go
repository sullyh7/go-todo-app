package model

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title string
	Text  string
	Done  bool
}
