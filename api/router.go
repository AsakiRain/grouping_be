package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/list", ListFile)
	r.GET("/read", ReadFile)
}
