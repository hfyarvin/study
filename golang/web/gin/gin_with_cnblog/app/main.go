package main

import (
	"../routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Router(r)
	r.Run(":8001")
}

func init() {

}