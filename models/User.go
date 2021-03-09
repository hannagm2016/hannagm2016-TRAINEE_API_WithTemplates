package models

type User struct{
    Id float64
    Name, UserName, Email, Phone, Website,Password string}

    func (p *PostModel) CreateUser(user User){
    p.db.Create(&user)
    }

func (p *PostModel) FindCustomerByName(userName string) User {
	var user User
	p.db.Find(&user, "user_name = ?", userName)
  	return user
}
func (p *PostModel) FindCustomerByEmail(email string) User {
	var user User
	p.db.Find(&user, "email = ?", email)
  	return user
}

func (p *PostModel) FindCustomerId(userName string) float64 {
	var id float64
	p.db.Table("users").Select("id").Find(&id, "user_name = ?", userName)
  	return id
}