package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router  ；路由
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})

}
