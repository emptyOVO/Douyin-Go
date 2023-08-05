package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5Encryption MD5加密封装
func Md5Encryption(password string) string {
	data := []byte(password)
	hash := md5.Sum(data)
	res := fmt.Sprintf("%x", hash)
	return res
}
