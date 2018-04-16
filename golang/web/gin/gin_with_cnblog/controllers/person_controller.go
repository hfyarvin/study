package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	p := &models.Person{LastName: "Wong", FirstName: "Arvin"}
	t := p.AddPerson()
	c.JSON(200, t)
}