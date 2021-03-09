package main

import (
    "fmt"
   "gorm.io/gorm"
   "gorm.io/driver/mysql"
   "encoding/json"
   "net/http"
   "io/ioutil"
   "site/models"

)

var db *gorm.DB

func init() {
var err error
 dsn := "mysql:mysql@tcp(127.0.0.1:3306)/articles?charset=utf8mb4&parseTime=True&loc=Local"
   db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
           if err != nil {
                panic("failed to connect database")
   }
}
//defer db.Close()

func takeUsers (url string) []models.User {
resp, err := http.Get(url)
          if err != nil {
             fmt.Println(err) }

         defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
 var users [] models.User
              errr := json.Unmarshal(body, &users)
                  if errr != nil {
                    fmt.Println (errr)
                 }
                 return users
          }

func main() {
db.AutoMigrate(&models.User{})
    users:=takeUsers("https://jsonplaceholder.typicode.com/users")
    for _,user:=range users {
          db.Create(&user)
    }
}
