package routers

import (
	"../controllers"
	"../controllers/test_controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &test_controllers.TestController{}, "get:Test")
}
