package util

import (
	"math/rand"

	"github.com/rs/xid"
)

// RandomStr 生成随机字符串
func RandomStr(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$@")
	var lenthLetter = len(letterRunes)

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(lenthLetter)]
	}
	return string(b)
}

// UUID 生成一个全局唯一的 id 字符串
// 实际上是改造后的 uuid.v4
func UUID() string {
	return xid.New().String()
}
