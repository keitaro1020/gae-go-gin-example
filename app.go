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

	b := r.Group("/books")
	{
		b.GET("", h.GetBooks)
		b.POST("", h.CreateBook)
		b.GET("/:id", h.GetBook)
		b.PUT("/:id", h.PutBook)
	}

	return r
}
