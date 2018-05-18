package handler

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func (h *Handler) Hello(c *gin.Context) {
	ac := appengine.NewContext(c.Request)

	log.Infof(ac, "hello!")
	c.String(http.StatusOK, "hello world")
}
