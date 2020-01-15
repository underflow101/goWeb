package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid Email or password"}
	}

	// Save User
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("bird").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("bird").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid Email or password"}
		}
		return
	
// JWT
// Create token
token := jwt.New(jwt.SigningMethodHS256)

// Set claim
claims := token.Claims.(jwt.MapClaims)
claims["id"] = u.ID
claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

