package ky3k

import "strconv"

func IntToString(d int) string {

	return strconv.Itoa(d)
}

func Int64ToString(d int64) string {

	return strconv.FormatInt(d, 10)
}
