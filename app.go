package app

import (
	"github.com/gin-gonic/gin"
	"github.com/keitaro1020/gae-go-gin-example/handler"
	"net/http"
)

func init() {
	http.Handle("/", router())
}

func router() http.Handler {
	h := handler.New()
	r := gin.New()

	r.GET("/hello", h.Hello)

	return r
}
