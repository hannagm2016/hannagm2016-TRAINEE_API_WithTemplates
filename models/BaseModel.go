package models

type BaseModel struct {
	IsAuthorized bool
	Posts        []Post
	Cust         User
}
