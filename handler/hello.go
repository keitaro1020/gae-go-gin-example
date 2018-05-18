package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}