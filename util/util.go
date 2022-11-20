package util

import (
	"regexp"
)

var dict = map[string]string{
	"姓名": "name",
	"年龄": "age",
	"性别": "sex",
	"学号": "id",
	"班级": "class",
	"年级": "grade",
	"学院": "college",
	"专业": "major",
	"序号": "serial",
}

func FindAlias(word string) string {
	reg := regexp.MustCompile(`\s+`)
	word = reg.ReplaceAllString(word, "")
	if alias, ok := dict[word]; ok {
		return alias
	}
	return word
}

func CleanRows(rows [][]string) [][]string {
	var result [][]string
	for _, row := range rows {
		if len(row) != 0 {
			result = append(result, row)
		}
	}
	// log.Println(result)
	return result
}
