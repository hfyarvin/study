package test_controller

import (
	"github.com/gin-gonic/gin"
)

func GoToThePage(c *gin.Context) {
	temp := []float64{1, 2, 5}
	hum := []float64{1, 2, 3}
	c_time := []string{"2012", "2013", "2014"}
	var list []Humiture
	for i := 0; i < 3; i++ {
		h := Humiture{Temp: temp, Hum: hum}
		list = append(list, h)
	}

	obj := gin.H{
		"list":  list,
		"temp":  temp,
		"ctime": c_time,
	}
	c.HTML(200, "index.html", obj)
}

type Humiture struct {
	Temp []float64
	Hum  []float64
}
