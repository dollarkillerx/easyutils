package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取sha1
func Sha1Encode(str string) string {
	data := []byte(str)
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}
