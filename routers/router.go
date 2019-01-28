package routers

import (
	"sms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/GetStationList", &controllers.SessionController{}, "get:GetStationList")
	beego.Router("/api/v1.0/GetUserList", &controllers.SessionController{}, "get:GetUserList")
}
