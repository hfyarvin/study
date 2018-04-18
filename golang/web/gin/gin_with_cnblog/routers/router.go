package routers

import (
	. "../controllers"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/hello", controllers.Hello)
	person := r.Group("/person")
	person.POST("/new", CreatePerson)

	r.LoadHTMLGlob("../views/*")

	home := r.Group("/home")
	{
		home.GET("/index", ShowHtmlPage)
		//list
		home.GET("/list", ListHtml)
		home.POST("/PageData", GetDataList)
		home.POST("/PageNextData", PageNextData)

		//add new page
		home.GET("/add", AddHtml)
		home.POST("/saveadd", AddPersonApi)

		//edit
		home.GET("/edit", EditHtml)
		home.POST("/saveedit", EditPersonApi)

		//del
		home.POST("/delete", DeletePersonApi)
	}
}
