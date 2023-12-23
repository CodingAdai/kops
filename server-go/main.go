// @Time    : 11/22/2023 9:27 PM
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!

package main

import (
	"server-go/config"
	"server-go/controller"
	"server-go/middle"
	"server-go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化k8s clientset
	service.K8s.Init()

	//初始化路由配置
	r := gin.Default()
	//跨域配置
	r.Use(middle.Cors())
	//jwt token验证
	//r.Use(middle.JWTAuth())
	//初始化路由
	controller.Router.InitApiRouter(r)

	r.StaticFile("/", "../dist/index.html")
	r.StaticFile("/index.html", "../dist/index.html")
	r.StaticFile("/app-loading.css", "../dist/app-loading.css")
	r.StaticFile("/app-loading2.css", "../dist/app-loading2.css")
	r.Static("/static", "../dist/static")

	//http server启动
	r.Run(config.ListenAddr)
}
