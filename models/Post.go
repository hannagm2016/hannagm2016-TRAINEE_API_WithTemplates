package models

import (
	"fmt"
	"gorm.io/gorm"
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

type PostModel struct {
	db *gorm.DB
}

type PostModelImpl interface {
	FindByID(id float64) Post
	FindAll() []Post
	DeleteByID(id float64)
	SaveByID(post Post)
	CreateUser(user User)
	FindCustomerByName(userName string) User
	FindCustomerByEmail(email string) User
	FindCustomerId(name string) float64
}

func NewPostModel(db *gorm.DB) *PostModel {
	return &PostModel{
		db: db,
	}
}

func (p *PostModel) FindAll() []Post {
	Posts := []Post{}
	p.db.Find(&Posts)
	return Posts
}
func (p *PostModel) FindByID(id float64) Post {
	post := Post{}
	p.db.Find(&post, "id = ?", id)
	p.db.Find(&post.Comment, "post_id = ?", id)
	return post
}
func (p *PostModel) DeleteByID(id float64) {
	var Posts Post
	p.db.Delete(&Posts, "id = ?", id)
	fmt.Println("***")
}
func (p *PostModel) SaveByID(post Post) {
	if post.Id != 0 {
		p.db.Save(&post)
	} else {
		p.db.Create(&post)
	}
	fmt.Println("**")
}
