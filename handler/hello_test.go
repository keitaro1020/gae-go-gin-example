package handler

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"github.com/favclip/testerator"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	instance, _, _ := testerator.SpinUp()
	defer testerator.SpinDown()

	h := New()
	r := gin.New()
	r.GET("/hello", h.Hello)
	req, _ := instance.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
