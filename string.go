package ky3k

import (
	"strconv"
)

func StringToInt(s string) (int, error) {

	return strconv.Atoi(s)
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}