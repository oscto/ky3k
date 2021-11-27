package ky3k

import (
	"crypto/md5"
	"encoding/hex"
)

func MarshalMd5(str, block string) string {

	m := md5.New()
	m.Write([]byte(str))
	m.Write([]byte(block))
	c := m.Sum(nil)

	return hex.EncodeToString(c)
}
