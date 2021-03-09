package main

import (
    "fmt"
   "gorm.io/gorm"
   "gorm.io/driver/mysql"
   "encoding/json"
   "net/http"
   "io/ioutil"
    "site/models"
"strconv"
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

func takePosts (url string) []models.Post {
resp, err := http.Get(url)
          if err != nil {
             fmt.Println(err) }

         defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
 var posts [] models.Post
              errr := json.Unmarshal(body, &posts)
                  if errr != nil {
                    fmt.Println (errr)
                 }
                 return posts
          }

func takeComments (url string) []models.Comments {
resp, err := http.Get(url)
          if err != nil {
             fmt.Println(err) }

         defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
 var comments [] models.Comments
              errr := json.Unmarshal(body, &comments)
                  if errr != nil {
                    fmt.Println (errr)
                 }
                 return comments
          }

func main() {
db.AutoMigrate(&models.Post{})
db.AutoMigrate(&models.Comments{})
    posts:=takePosts("https://jsonplaceholder.typicode.com/posts?userId=7")
    for _,post:=range posts {
          db.Create(&post)
             comm:=takeComments("https://jsonplaceholder.typicode.com/comments?postId="+strconv.Itoa(int(post.Id)))
                  for _,comment:=range comm {
                      db.Create(&comment)
                  }
    }
}
