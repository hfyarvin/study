package controllers

import (
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