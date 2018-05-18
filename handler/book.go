package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/keitaro1020/gae-go-gin-example/domain"
	"github.com/keitaro1020/gae-go-gin-example/repository"
	"google.golang.org/appengine"
	"net/http"
)

func (h *Handler) CreateBook(c *gin.Context) {
	ac := appengine.NewContext(c.Request)

	b := domain.NewBook(CreateNewId())
	if err := c.Bind(b); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	r := repository.NewBookRepository()
	b, err := r.PutBook(ac, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, b)
}

func (h *Handler) GetBooks(c *gin.Context) {
	ac := appengine.NewContext(c.Request)

	r := repository.NewBookRepository()
	bs, err := r.GetBooks(ac)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bs)
}
