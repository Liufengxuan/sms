package controllers

import (
	"sms/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type SessionController struct {
	beego.Controller
}

func (ctx *SessionController) retData(resp map[string]interface{}) {
	ctx.Data["json"] = resp
	beego.Debug(resp)
	ctx.ServeJSON()
}

func (ctx *SessionController) GetStationList() {
	resp := make(map[string]interface{})
	defer ctx.retData(resp)
	ini, err := config.NewConfig("ini", "conf/db.conf")

	if err != nil {
		resp["state"] = models.RECODE_INI_ReadERR
		resp["msg"] = err
		return
	}
	temp := ini.String("dblist::database")
	beego.Debug(temp)
	if len(temp) < 2 {
		resp["state"] = models.RECODE_INI_ReadERR
		resp["msg"] = models.RecodeText(models.RECODE_INI_ReadERR)
		return
	}

	dblist := strings.Split(temp, "&")
	beego.Debug(dblist)
	var stationList []models.Station
	for _, v := range dblist {
		t := strings.Split(v, ":")
		if len(t) == 2 {
			a := models.Station{t[0], t[1]}
			stationList = append(stationList, a)
		}
	}
	if len(stationList) > 0 {
		resp["data"] = stationList
		resp["state"] = models.RECODE_OK
		resp["msg"] = models.RecodeText(models.RECODE_OK)
		return
	} else {
		resp["state"] = models.RECODE_INI_ReadERR
		resp["msg"] = "/conf/db.conf数据库信息未能正确解析。"
	}
}
