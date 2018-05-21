package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mercari.io/datastore"
	"net/http"
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
