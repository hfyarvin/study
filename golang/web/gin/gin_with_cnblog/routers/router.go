package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine)  {
	r.GET("/hello", controllers.Hello)
	person := r.Group("/person")
	person.POST("/new", controllers.CreatePerson)
}