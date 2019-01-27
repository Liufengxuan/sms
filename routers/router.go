package routers

import (
	"sms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/GetStationList", &controllers.SessionController{}, "get:GetStationList")
}
