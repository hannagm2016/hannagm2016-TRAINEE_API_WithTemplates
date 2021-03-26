package models

import (
//"gorm.io/gorm"
)

type Post struct {
	UserId  float64
	Id      float64
	Title   string
	Body    string
	Comment []Comments `gorm:"-"`
}
type Comments struct {
	Name   string
	Email  string
	Body   string
	Id     float64
	PostId float64
}
