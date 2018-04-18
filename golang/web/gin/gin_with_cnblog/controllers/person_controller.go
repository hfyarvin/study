package controllers

import (
	"net/http"
	"strconv"

	"../models"
	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	lastName := c.Query("last_name")
	firstName := c.Query("first_name")
	p := &models.Person{LastName: lastName, FirstName: firstName}
	t := p.AddPerson()
	c.JSON(200, t)
}

//渲染html
func ShowHtmlPage(c *gin.Context) {
	obj := gin.H{
		"title": "GIN: HTML页面",
	}
	c.HTML(http.StatusOK, "index.html", obj)
}

//List
func ListHtml(c *gin.Context) {
	obj := gin.H{
		"title": "GIN: 用户列表页面",
	}
	c.HTML(http.StatusOK, "list.html", obj)
}

//用户列表
func GetDataList(c *gin.Context) {
	search := c.PostForm("search")
	num := c.PostForm("pageno")
	pageno, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	dataList := models.GetPersonList(pageno, 3, search)

	count := models.GetRecordNum(search)

	obj := gin.H{
		"list":     dataList,
		"count":    count,
		"pagesize": 3,
		"pageno":   pageno,
	}

	c.JSON(http.StatusOK, obj)
}

//列表页面数据
func PageNextData(c *gin.Context) {
	//得到请求的参数
	search := c.PostForm("search")
	num := c.PostForm("pageno")
	pageno, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	//得到数据集
	datalist := GetPersonList(pageno, 3, search)

	//得到记录的总数
	count := GetRecordNum(search)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"datalist": datalist,
		"count":    count,
		"pagesize": 3,
		"pageno":   pageno,
	})
}

//新增页面
func AddHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "GIN: 新增用户页面",
	})
}

//新增记录
func AddPersonApi(c *gin.Context) {

	//得到请求的参数
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	fmt.Println("执行到此处A")
	//赋值
	p := Person{FirstName: firstName, LastName: lastName}
	//调用模型中的新增的方法
	ra := p.AddPerson()
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": ra,
	})

	c.Redirect(http.StatusOK, "/home/list")
}

//编辑页面
func EditHtml(c *gin.Context) {
	//得到URL请求的参数
	num := c.Query("id")

	id, err := strconv.Atoi(num)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	p := models.GetPersonById(id)
	if p == nil {
		fmt.Println("得到数据错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "得到数据错误",
		})
		c.Abort()
		return
	} else {
		fmt.Println(p)
		fmt.Println("得到数据正确")
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"title":     "GIN: 编辑用户页面",
		"id":        p.Id,
		"firstname": p.FirstName,
		"lastname":  p.LastName,
	})
}

//编辑记录
func EditPersonApi(c *gin.Context) {
	fmt.Println("执行到此处")
	//得到请求的参数
	num := c.PostForm("id")
	fmt.Println(num)
	id, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	//赋值
	p := models.GetPersonById(id)
	p.FirstName = firstName
	p.LastName = lastName
	//调用模型中的编辑的方法
	ra := p.EditPerson()
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": ra,
	})
	//跳转Home
	c.Redirect(http.StatusOK, "/home/list")
}

//删除记录
func DeletePersonApi(c *gin.Context) {

	//得到请求的参数
	num := c.PostForm("id")
	fmt.Println(num)
	id, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	//调用模型中的删除的方法
	ra := DeletePerson(id)
	if ra == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "删除失败",
		})
		c.Abort()
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": ra,
	})
}
