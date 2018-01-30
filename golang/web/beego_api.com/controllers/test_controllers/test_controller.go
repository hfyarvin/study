package test_controllers

import (
	"../../controllers"
	"../../models/user_model"
)

type TestController struct {
	controllers.MainController
}

func (self *TestController) Test() {
	num := user_model.GetDataNum()
	self.Data["json"] = num
	self.ServeJSON()
}
