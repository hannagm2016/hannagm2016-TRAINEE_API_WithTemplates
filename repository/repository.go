package repository

import (
	"fmt"
	"gorm.io/gorm"
	"site/models"
)

type PostModel struct {
	db *gorm.DB
}

type PostModelImpl interface {
	FindByID(id float64) models.Post
	FindAll() []models.Post
	DeleteByID(id float64)
	SaveByID(post models.Post)
	CreateUser(user models.User)
	FindCustomerByName(userName string) models.User
	FindCustomerByEmail(email string) models.User
	FindCustomerId(name string) float64
}

func NewPostModel(db *gorm.DB) *PostModel {
	return &PostModel{
		db: db,
	}
}

func (p *PostModel) FindAll() []models.Post {
	Posts := []models.Post{}
	p.db.Find(&Posts)
	return Posts
}
func (p *PostModel) FindByID(id float64) models.Post {
	post := models.Post{}
	p.db.Find(&post, "id = ?", id)
	p.db.Find(&post.Comment, "post_id = ?", id)
	return post
}
func (p *PostModel) DeleteByID(id float64) {
	var Posts models.Post
	p.db.Delete(&Posts, "id = ?", id)
	fmt.Println("***")
}
func (p *PostModel) SaveByID(post models.Post) {
	if post.Id != 0 {
		p.db.Save(&post)
			fmt.Println(post.Id, "000")

	} else {
	fmt.Println(post.Id, "!!!!000")
		p.db.Create(&post)
	}
	fmt.Println("**")
}
