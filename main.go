package main

import (
	"embed"
	"grouping_be/api"
	"grouping_be/middleware"
	"grouping_be/static"
	"grouping_be/util"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed dist
	dist embed.FS
)

func main() {
	r := gin.Default()
	//配置跨域（前后端分离调试用）
	r.Use(middleware.Cors())
	//服务前端静态文件
	r.Use(static.Serve("/", static.EmbedFolder(dist, "dist")))
	//服务后端api
	api.SetupRouter(r)
	//打开浏览器
	go util.OpenBrowser("http://localhost:12870")
	//选择高端口启动服务器
	r.Run("0.0.0.0:12870") // 0.0.0.0:8080
}
