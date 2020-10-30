package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router  ；路由
    beego.Router("/", &controllers.MainController{})
	//用户注册接口
    beego.Router("/register", &controllers.RegisterController{})
	//用户登录接口
    beego.Router("/login", &controllers.LoginController{})
	//错误页面接口
    beego.Router("/error", &controllers.Error{})
	//用户上传的功能
	beego.Router("/upload",&controllers.UploadFileController{})
    //查看认证数据证书
    beego.Router("/cert_detail.html",&controllers.CertDataController{})
	//用户实名认证请求
	beego.Router("/user_kyc.html",&controllers.UserKycController{})

}
