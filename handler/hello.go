package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
)

func (h *Handler) Hello(c *gin.Context) {
	ac := appengine.NewContext(c.Request)

	log.Infof(ac, "hello!")
	c.String(http.StatusOK, "hello world")
}