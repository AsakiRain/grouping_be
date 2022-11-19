package api

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func ListFile(ctx *gin.Context) {
	files, err := os.ReadDir(".")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": err.Error(),
			"data":    nil,
		})
	}

	var fileList []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		fileExt := path.Ext(fileName)
		if fileExt == ".xlsx" {
			fileList = append(fileList, fileName)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "success",
		"data": map[string]interface{}{
			"files": fileList,
		},
	})
}
