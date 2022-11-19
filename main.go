package main

import (
	"embed"
	"grouping_be/api"
	"grouping_be/middleware"

	"grouping_be/static"

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
	//选择高端口启动服务器
	r.Run("0.0.0.0:12870") // 0.0.0.0:8080
}
