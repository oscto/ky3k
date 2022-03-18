package ky3k

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
)

func MarshalMd5(str, block string) string {

	m := md5.New()
	m.Write([]byte(str))
	m.Write([]byte(block))
	c := m.Sum(nil)

	return hex.EncodeToString(c)
}

func MarshalSHA1(str string) string {

	sha := sha1.New()
	sha.Write([]byte(str))
	c := sha.Sum(nil)

	return hex.EncodeToString(c)
}

func MarshalCRC32(str string) uint32 {
	
	return crc32.ChecksumIEEE([]byte(str))
}
