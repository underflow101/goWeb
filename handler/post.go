package handler

import (
	"goWeb/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	u := &model.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	p := &model.Post{
		ID:   bson.NewObjectId(),
		From: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.To == "" || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid target or message"}
	}

	// Find user from db
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in db
	if err = db.DB("bird").C("posts").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Default
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retreive post from db
	posts := []*model.Post{}
	db := h.DB.Clone()
	if err = db.DB("bird").C("posts").
		Find(bson.M{"to": userID}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, posts)
}
