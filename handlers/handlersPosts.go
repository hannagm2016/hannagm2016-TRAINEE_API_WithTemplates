package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"site/models"
	"strconv"
	"site/repository"
)

const (
	COOKIE_NAME = "sessionId"
)

var Posts []models.Post

type handler struct {
	PostModel repository.PostModelImpl
}

func NewHandler(p repository.PostModelImpl) *handler {
	return &handler{p}
}

// Index godoc
// @Summary Show all post
// @Description all posts
// @Tags Posts
// @Produce  json
// @Success 200 {array} models.Post
// @Failure 500 {string} string "fail"
// @Router / [get]
func (h *handler) Index(c echo.Context) error {
	Posts = h.PostModel.FindAll()
	var Model models.BaseModel
	cookie, _ = c.Cookie(COOKIE_NAME)
	if cookie != nil {
		Model.IsAuthorized = true
	}
	Model.Posts = Posts
	Model.Cust = Customer
    	  fmt.Println("Endpoint Hit: Index" )
	return c.JSON(http.StatusOK, Model)
}

// ReturnSinglePost godoc
// @Summary Show single post with id specified
// @Description get post by ID
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param id path float64 true "post id"
// @Success 200 {object} models.Post
// @Failure 404 {string} string "not found"
// @Router /post/{id} [get]
func (h *handler) ReturnSinglePost(c echo.Context) error {
	id := c.Param("id")
	key, _ := strconv.ParseFloat(string(id), 64)
	post := h.PostModel.FindByID(key)
	return c.Render(http.StatusOK, "onepost2.html", post)
}

// DeletePost godoc
// @Summary Delete one post with id specified
// @Description delete post by ID
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param id path float64 true "post id"
// @Success 200 {string} string "success"
// @Failure 403 {string} string "not registered"
// @Failure 404 {string} string "not found"
// @Router /deletePost/{id} [delete]
func (h *handler) DeletePost(c echo.Context) error {
	cookie, _ = c.Cookie(COOKIE_NAME)
	if cookie == nil {
		//c.String(http.StatusForbidden, "Not registered")
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	id := c.Param("id")
	key, _ := strconv.ParseFloat(string(id), 64)
	h.PostModel.DeleteByID(key)
	fmt.Println("Endpoint Hit: DeletePost", id)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// SavePost godoc
// @Summary Save new or updated post
// @Description Save post to db
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param body body {object} models.Post true "post body"
//@Success 200 {string} string "success - redirect to index"
// @Failure 500 {string} string "fail"
// @Router /savePost [post]
func (h *handler) SavePost(c echo.Context) error {
	var post models.Post
	post.Id, _ = strconv.ParseFloat(c.FormValue("id"), 0)
	post.Title = c.FormValue("title")
	post.Body = c.FormValue("body")
	post.UserId = Customer.Id //Later would be User ID from authorisation
	h.PostModel.SaveByID(post)

	fmt.Println("Endpoint Hit: InsertrPost")
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// CreateNewPost godoc
// @Summary Form for creation new post
// @Description Form for creation new post
// @Tags Posts
// @Accept  json
// @Produce  xml
// @Param id path int true "Post ID"
//@Success 200 {string} string "success - redirect to save post"
// @Failure 403 {string} string "not registered"
// @Failure 404 {string} string "not found"
// @Router /post [get]
func (h *handler) CreateNewPost(c echo.Context) error {
	cookie, _ = c.Cookie(COOKIE_NAME)
	fmt.Println(cookie)
	if cookie == nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{})
}

// EditPost godoc
// @Summary Form for updating post
// @Description Get post and edit it
// @Tags Posts
// @Accept  json
// @Produce  json
// @Param id path float64 true "post id"
//@Success 200 {string} string "success - redirect to save post"
// @Failure 403 {string} string "not registered"
// @Failure 404 {string} string "not found"
// @Router /postUpdate/{id} [get]
func (h *handler) EditPost(c echo.Context) error {
	cookie, _ = c.Cookie(COOKIE_NAME)
	if cookie == nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
	fmt.Println(Customer)
	id := c.Param("id")
	key, _ := strconv.ParseFloat(string(id), 64)
	post := h.PostModel.FindByID(key)
	return c.Render(http.StatusOK, "create.html", post)

}
