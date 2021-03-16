package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"site/models"
	"site/session"
	"strconv"
	"strings"
	"time"
	"site/config"
)

var inMemorySession *session.Session

var Customer models.User

var cookie *http.Cookie

type AccessToken struct {
	Token  string
	Expiry int64
}

func readHttpBody(response *http.Response) string {
	fmt.Println("Reading body")
	bodyBuffer := make([]byte, 5000)
	var str string
	count, err := response.Body.Read(bodyBuffer)
	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
		}
		str += string(bodyBuffer[:count])
	}
	fmt.Println(str)
	return str
}

func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
	fmt.Println("GetAccessToken")
	var token AccessToken
	if callbackUri == "http://localhost:8000/FBLogin" { //For Facebook
		response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
			client_id + "&redirect_uri=" + callbackUri +
			"&client_secret=" + secret + "&code=" + code)
		if err == nil {
			auth := readHttpBody(response)
			tokenArr := strings.Split(auth, ",")
			token.Token = strings.Split(tokenArr[0], ":")[1]
			token.Token = token.Token[1 : len(token.Token)-1]
			expire := strings.Split(tokenArr[2], ":")[1]
			expireInt, err := strconv.Atoi(strings.Split(expire, "}")[0])
			if err == nil {
				token.Expiry = int64(expireInt)
			}
		}
	} else { //For Google
		postBody, _ := json.Marshal(map[string]string{
			"code":          code,
			"client_id":     client_id,
			"client_secret": secret,
			"redirect_uri":  callbackUri,
			"grant_type":    "authorization_code",
		})
		responseBody := bytes.NewBuffer(postBody)
		response, err := http.Post("https://accounts.google.com/o/oauth2/token", "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer response.Body.Close()
		//	if err == nil {
		auth := readHttpBody(response)
		tokenArr := strings.Split(auth, ",")
		token.Token = strings.Split(tokenArr[0], ":")[1]
		token.Token = token.Token[1 : len(token.Token)-1]
		expire := strings.Split(tokenArr[1], ":")[1]
		expireInt, err := strconv.Atoi(expire)
		if err == nil {
			token.Expiry = int64(expireInt)
		}
	}
	return token
}

// Authorisation godoc
// @Summary Authorisation form
// @Description Entering login and pass into authorisation form
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "authorisation form"
// @Failure 500 {string} string "fail"
// @Router /authorisation [get]
func (h *handler) Authorisation(c echo.Context) error {

	inMemorySession = session.NewSession()
	fmt.Println("Endpoint Hit: authorisation")
fb:= config.GetConfig("fb")
	fbConfig := &oauth2.Config{
	    ClientID: fb.ClientID,
		ClientSecret: fb.ClientSecret,
		RedirectURL:  fb.RedirectURL,
		Scopes:       []string{"email", "user_birthday", "user_location"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	FB := fbConfig.AuthCodeURL("")
	google:=config.GetConfig("google")
	GoogleConfig := oauth2.Config{
		ClientID: google.ClientID,
        ClientSecret: google.ClientSecret,
        RedirectURL:  google.RedirectURL,
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	Google := GoogleConfig.AuthCodeURL("")
fmt.Println(FB, Google)
	return c.Render(http.StatusOK, "authorisation.html", map[string]string{
		"FB":     FB,
		"Google": Google,
	})
}

func (h *handler) FBLogin(c echo.Context) error {
	code := c.FormValue("code")
	fb:= config.GetConfig("fb")
	accessToken := GetAccessToken(fb.ClientID, code, fb.ClientSecret, fb.RedirectURL)
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.Token)
	if err != nil {
		fmt.Println(err)
	}
	str := readHttpBody(response)
	user, _ := jason.NewObjectFromBytes([]byte(str))
	name, _ := user.GetString("name")
	sessionId := inMemorySession.Init(name)
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	c.SetCookie(cookie)
	Customer = h.PostModel.FindCustomerByName(name)
	if Customer.UserName == "" {
		Customer.Name = name
		Customer.UserName = name
		Customer.Website = "FB"
		h.PostModel.CreateUser(Customer)
		Customer.Id = h.PostModel.FindCustomerId(name)
	}
	fmt.Println(Customer)
	fmt.Println("Endpoint Hit: FBLogin", cookie)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
func (h *handler) GoogleLogin(c echo.Context) error {
	code := c.FormValue("code")
	google:=config.GetConfig("google")
	accessToken := GetAccessToken(google.ClientID, code, google.ClientSecret, google.RedirectURL)
	response, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + accessToken.Token)
	if err != nil {
		fmt.Println(err)
	}
	str := readHttpBody(response)
	user, _ := jason.NewObjectFromBytes([]byte(str))
	email, _ := user.GetString("email")
	sessionId := inMemorySession.Init(email)
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	c.SetCookie(cookie)
	Customer = h.PostModel.FindCustomerByName(email)
	if Customer.UserName == "" {
		Customer.Email = email
		Customer.UserName = email
		Customer.Website = "Google"
		h.PostModel.CreateUser(Customer)
		Customer.Id = h.PostModel.FindCustomerId(email)
	}
	fmt.Println("Endpoint Hit: GoogleLogin", cookie)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// AuthorisationPost godoc
// @Summary Authorisation process
// @Description Creating cookie, set IsAuthorize as true
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "user authorised"
// @Failure 500 {string} string "fail"
// @Router /authorisationPost [post]
func (h *handler) AuthorisationPost(c echo.Context) error {
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword")
	fmt.Println(inputPassword, inputEmail)
	sessionId := inMemorySession.Init(inputEmail)
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	Customer = h.PostModel.FindCustomerByEmail(inputEmail)

	c.SetCookie(cookie)
	fmt.Println("Endpoint Hit: authorisation", Customer)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Registration godoc
// @Summary Registration form
// @Description Creating user, cookie, set IsAuthorize as true
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "redirect to registrationPost"
// @Failure 500 {string} string "fail"
// @Router /registration [get]
func (h *handler) Registration(c echo.Context) error {
	fmt.Println("Endpoint Hit: registration")
	return c.Render(http.StatusOK, "registration.html", map[string]interface{}{})
}

// RegistrationPost godoc
// @Summary Registration process
// @Description Creating user, cookie, set IsAuthorize as true
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "user registrate "
// @Failure 500 {string} string "fail"
// @Router /registrationPost [post]
func (h *handler) RegistrationPost(c echo.Context) error {

	Customer.Name = c.FormValue("inputName")
	Customer.Email = c.FormValue("inputEmail")
	Customer.Password = c.FormValue("inputPassword")
	Customer.UserName = c.FormValue("inputName")
	h.PostModel.CreateUser(Customer)
	Customer.Id = h.PostModel.FindCustomerId(Customer.UserName)
	sessionId := inMemorySession.Init(Customer.Email)
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
		MaxAge:  60 * 60,
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Logout godoc
// @Summary Log out, remove authorisation
// @Description Removing cookie, set IsAuthorize as false
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} string "user logged out"
// @Failure 500 {string} string "fail"
// @Router /logout [get]
func (h *handler) Logout(c echo.Context) error {
	fmt.Println("Endpoint Hit: logout", cookie)
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   "",
		Expires: time.Now(),
	}
	c.SetCookie(cookie)
	//  Customer=new(models.User)
	fmt.Println("Endpoint Hit: logout", cookie)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
