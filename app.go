package app

import (
	"github.com/gin-gonic/gin"
	"github.com/keitaro1020/gae-go-gin-example/handler"
	_ "go.mercari.io/datastore/aedatastore"
	"net/http"
)

func init() {
	http.Handle("/", router())
}

func router() http.Handler {
	h := handler.New()
	r := gin.New()

	r.GET("/hello", h.Hello)
	r.POST("/book", h.CreateBook)
	r.GET("/book", h.GetBooks)

	return r
}
