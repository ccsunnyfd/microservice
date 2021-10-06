package code

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func GenToken() string {
	return GenMD5(RandCode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 32))
}

func GenMD5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func RandCode(s string, size int) []byte {
	length := len(s)
	var res []byte
	for i := 0; i < size; i++ {
		pos := rand.Intn(length)
		res = append(res, s[pos])
	}
	return res
}
