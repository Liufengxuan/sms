package controllers

import (
	"fmt"
	"sms/models"
	"sms/models/status"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/denisenkom/go-mssqldb"
)

//SessionController 处理登陆前请求和登录请求。
type SessionController struct {
	beego.Controller
}

func (ctx *SessionController) retData(resp map[string]interface{}) {
	ctx.Data["json"] = resp
	beego.Debug("Response :" + fmt.Sprint(ctx.Data["json"]))
	ctx.ServeJSON()
}

//GetStationList 获取各个数据库的名称
func (ctx *SessionController) GetStationList() {
	resp := make(map[string]interface{})
	defer ctx.retData(resp)
	ini, err := config.NewConfig("ini", "conf/db.conf")

	if err != nil {
		resp["state"] = status.RECODE_INI_READERR
		resp["msg"] = err
		return
	}
	temp := ini.String("dblist::database")
	beego.Debug(temp)
	if len(temp) < 2 {
		resp["state"] = status.RECODE_INI_READERR
		resp["msg"] = status.RecodeText(status.RECODE_INI_READERR)
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
		resp["state"] = status.RECODE_OK
		resp["msg"] = status.RecodeText(status.RECODE_OK)
		return
	} else {
		resp["state"] = status.RECODE_INI_READERR
		resp["msg"] = "/conf/db.conf数据库信息未能正确解析。"
	}
}

//GetUserList 根据数据库的名称加载此数据库的用户列表。
func (ctx *SessionController) GetUserList() {
	resp := make(map[string]interface{})
	defer ctx.retData(resp)

	dbName := ctx.GetString("station", "")

	db, err := models.GetDB(dbName)
	if err != nil {
		beego.Error(err)
		resp["state"] = status.RECODE_DB_CONNERR
		resp["msg"] = status.RecodeText(status.RECODE_DB_CONNERR)
		return
	}
	rows, err := db.Query("select ID,PersonName as name,Pword as pwd,Dept,Duty   from TblLoginPerson")
	if err != nil {
		beego.Error(err)
		resp["state"] = status.RECODE_DB_QUERYERR
		resp["msg"] = status.RecodeText(status.RECODE_DB_QUERYERR)
		return
	}
	if rows == nil {
		resp["state"] = status.RECODE_DB_QUERYNUll
		resp["msg"] = status.RecodeText(status.RECODE_DB_QUERYNUll)
		return
	}
	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Pwd, &u.Dept, &u.Duty)
		users = append(users, u)
	}
	resp["state"] = status.RECODE_OK
	resp["msg"] = status.RecodeText(status.RECODE_OK)
	resp["data"] = users
	//ctx.Ctx.Request.Body.Read()
}
