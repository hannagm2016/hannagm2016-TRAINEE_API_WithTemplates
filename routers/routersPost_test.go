package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"site/models"
)

type PostsModelStub struct{}

func (y *PostsModelStub) FindAll() []models.Post {
	posts := []models.Post{}
	posts = append(posts, models.Post{
		UserId: 1.,
		Id:     1.,
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	})
	return posts
}

func TestGetIndex(t *testing.T) {
	e := echo.New()
	render_htmls := NewTemplate()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")
	e.Renderer = render_htmls

	p := &PostsModelStub{}
	h := NewHandler(p)

	var userJSON = `{"posts":[{
                                "UserId": 1.,
                                "Id": 1.,
                                "Title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
                                "Body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
                              }]}`

	if assert.NoError(t, h.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
