package repository
import (
    "site/models"
)

func (p *PostModel) CreateUser(user models.User) {
	p.db.Create(&user)
}

func (p *PostModel) FindCustomerByName(userName string) models.User {
	var user models.User
	p.db.Find(&user, "user_name = ?", userName)
	return user
}
func (p *PostModel) FindCustomerByEmail(email string) models.User {
	var user models.User
	p.db.Find(&user, "email = ?", email)
	return user
}

func (p *PostModel) FindCustomerId(userName string) float64 {
	var id float64
	p.db.Table("users").Select("id").Find(&id, "user_name = ?", userName)
	return id
}
