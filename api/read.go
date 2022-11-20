package api

import (
	"fmt"
	"grouping_be/util"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Field struct {
	Title     string `json:"title"`
	DataIndex string `json:"dataIndex"`
}

type Record map[string]interface{}

func ReadFile(ctx *gin.Context) {
	fileName := ctx.Query("filename")
	if fileName == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "没有提供文件名",
			"data":    nil,
		})
		return
	}

	ext := path.Ext(fileName)
	if ext != ".xlsx" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "不支持的文件格式",
			"data":    nil,
		})
		return
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": "文件不存在",
				"data":    nil,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": err.Error(),
				"data":    nil,
			})
		}
		return
	}

	file, err := excelize.OpenFile(fileName)
	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	sheets := file.GetSheetList()
	if len(sheets) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "文件中没有工作表",
			"data":    nil,
		})
		return
	}
	// log.Printf("文件中有 %d 个工作表：%v", len(sheets), sheets)

	rows, err := file.GetRows(sheets[0])
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	rows = util.CleanRows(rows)

	var fields []Field
	for _, field := range rows[0] {
		fields = append(fields, Field{
			Title:     field,
			DataIndex: util.FindAlias(field),
		})
	}

	var records []Record
	for _, row := range rows[1:] {
		record := make(Record)
		for i, value := range row {
			record[util.FindAlias(rows[0][i])] = value
		}
		records = append(records, record)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "success",
		"data": map[string]interface{}{
			"filename": fileInfo.Name(),
			"filesize": fileInfo.Size(),
			"modified": fileInfo.ModTime().Format("2006-01-02 15:04:05"),
			"ext":      ext,
			"sheet":    sheets[0],
			"colcount": len(fields),
			"rowcount": len(records),
			"fields":   fields,
			"records":  records,
		},
	})
}
