package ky3k

import (
	"crypto/md5"
	"encoding/hex"
)

func MarshalMd5(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte(salt))
	st := h.Sum(nil)

	return hex.EncodeToString(st)
}
