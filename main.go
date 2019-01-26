package main

import (
	"fmt"
	"net/http"
	_ "sms/routers"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	ignoreStaticPath()

	beego.Run()
}

func ignoreStaticPath() {
	//
	//请求url中没有api字段。
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

//TransparentStatic da
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path      //拿到请求的路径
	beego.Debug("request url:", orpath) //打印路径

	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}

	beego.Debug(fmt.Sprintf("page" + ctx.Request.URL.Path))                      //实际路径
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "page"+ctx.Request.URL.Path) //返回静态页面
}
