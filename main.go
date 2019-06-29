package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := newPath()
}

func addMiddleware(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		c.Abort()
		return
	}

	c.Next()
}

func newPath() *gin.Context {
	r := gin.Default()

	r.Use(addMiddleware)

	r.POST("/customers", service.PostCustomer)

	return r
}
