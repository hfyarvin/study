/**
 * 1. 获取form参数(application/x-www-form-urlencoded): c.PostForm
 **/
package router

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//单个路由的中间件
	r.POST("/user", Middelware(), UserInfo)
	r.GET("/user/detail", UserDetail)
	r.GET("/search", Baidu)
	r.POST("/upload", Upload)
	r.POST("/upload/all", UploadAll)
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

//上传文件"Content-Type: multipart/form-data"
func Upload(c *gin.Context) {
	// file, header, _ := c.Request.FormFile("file")
	header, _ := c.FormFile("file")
	err := saveFile(header)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	c.String(http.StatusOK, header.Filename+" uploaded!")
}

func UploadAll(c *gin.Context) {
	//可能是限制大小
	err := c.Request.ParseMultipartForm(20000)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	formdata := c.Request.MultipartForm //得到文件句柄
	files := formdata.File["upload"]
	for _, item := range files {
		err := saveFile(item)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
	}
	c.String(http.StatusOK, "Successful uploaded!")
}

func saveFile(header *multipart.FileHeader) error {
	file, err := header.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	filename := header.Filename
	out, createFileErr2 := os.Create(fmt.Sprintf("./static/file/%s", filename))
	defer out.Close()
	if createFileErr2 != nil {
		return createFileErr2
	}
	_, cErr := io.Copy(out, file)
	if cErr != nil {
		return cErr
	}
	return nil
}
