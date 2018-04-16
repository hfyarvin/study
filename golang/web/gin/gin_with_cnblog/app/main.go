package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run(":8001")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
