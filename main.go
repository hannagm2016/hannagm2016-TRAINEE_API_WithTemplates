package main

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	"html/template"
	"io"
	"site/db"
	_ "site/docs"
//	"site/models"
	"site/handlers"
	"site/repository"
	"github.com/labstack/echo/middleware"
	//"net/http"

)

// @title Swagger Example API for trainee exercise
// @version 1.0
// @description This is an implementation of api server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {

	render_htmls := NewTemplate()
	render_htmls.Add("authorisation.html", template.Must(template.ParseFiles("templates/authorisation.html")))
	render_htmls.Add("registration.html", template.Must(template.ParseFiles("templates/registration.html")))
	e := echo.New()

	e.Use(middleware.CORS())
	d := db.DBConnect()
	h := handlers.NewHandler(repository.NewPostModel(d))
	e.Renderer = render_htmls
	e.File("/", "templates/index2.html")
	e.GET("/posts", h.Index)
	e.GET("/posts/:id", h.ReturnSinglePost)//was GET
	e.GET("/authorization", h.Authorisation)
	e.GET("/logout", h.Logout)
	e.GET("/registration", h.Registration)
	e.GET("/FBLogin", h.FBLogin)
	e.GET("/GoogleLogin", h.GoogleLogin)
	e.DELETE("/deletePost/:id", h.DeletePost)
	e.POST("/authorisationPost", h.AuthorisationPost)
	e.POST("/registrationPost", h.RegistrationPost)
	e.POST("/savePost", h.SavePost)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))

}

const (
	COOKIE_NAME = "sessionId"
)

type Template struct { //the map[key] in key means 'Your html file name'
	templates map[string]*template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: make(map[string]*template.Template),
	}
}

func (t *Template) Render(w io.Writer, html_name string, data interface{}, c echo.Context) error {

	if tmpl, exist := t.templates[html_name]; exist { //Check existence of the t.templates[html_name]
		return tmpl.Execute(w, data) // ** It wll execute the map[string]interface{} data
	} else {
		return errors.New("There is no " + html_name + " in Template map.")
	}
}
func (tmpl *Template) Add(html_name string, template *template.Template) {
	tmpl.templates[html_name] = template
}
