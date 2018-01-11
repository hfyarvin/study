package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//单个路由的中间件
	r.POST("/user", Middelware(), UserInfo)
	r.GET("/user/detail", UserDetail)
	r.GET("/search", Baidu)
}

//数据绑定解析
type User struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

func UserInfo(c *gin.Context) {
	u := new(User)
	//绑定数据
	err := c.BindJSON(u)
	if err == nil {
		c.JSON(http.StatusOK, u)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

//重定向
func Baidu(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}

//自定义中间件
func Middelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 在gin上下文中定义变量
		c.Set("example", "12345")

		// 请求前

		c.Next() //处理请求

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

//绑定query
func UserDetail(c *gin.Context) {
	type U struct {
		Email    string `form:"email"`
		UserName string `form:"user_name"`
	}
	u := new(U)
	err := c.ShouldBindQuery(u)
	if err == nil {
		c.JSON(200, u)
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
