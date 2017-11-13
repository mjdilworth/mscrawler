package main

import (

	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestDefaultLanding(t *testing.T) {


	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
	})

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		// If this fails, it's because httprouter needs to be updated to at least f78f58a0db
		t.Errorf("Status code should be %v, was %d. Location: %s", http.StatusNotFound, w.Code, w.HeaderMap.Get("Location"))
	}
}

/*
func DefaultLanding(c *gin.Context) {
	c.String(http.StatusOK, "mjd test")
}
*/
