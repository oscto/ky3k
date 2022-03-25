package ky3k

import (
	"strconv"
	"unicode"
)

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for _, v := range str {
		f := string(unicode.ToUpper(v))
		return f + str[len(f):]
	}

	return ""
}
