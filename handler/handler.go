package handler

import (
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mercari.io/datastore"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func CreateNewId() string {
	return uuid.New().String()
}

func ErrorResponse(c *gin.Context, err error) {
	ec := http.StatusInternalServerError

	if err == datastore.ErrNoSuchEntity {
		ec = http.StatusNotFound
	}

	c.JSON(ec, err)
}