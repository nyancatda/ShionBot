package utils

import (
	"strconv"
	"strings"
)

//替换字符串中的变量位置
func StringVariable(ReturnText string, text []string) string {
	for value := 0; value < len(text); value++ {
		ReturnText = strings.Replace(ReturnText, "{%"+strconv.Itoa(value)+"}", text[value], 1)
	}

	return ReturnText
}
