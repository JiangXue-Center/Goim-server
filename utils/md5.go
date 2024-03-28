package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

func ValidPassword(plainpwd, salt, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}

// generateSalt 生成指定长度的随机密码盐值
func GenerateSalt(length int) (string, error) {
	// 创建一个指定长度的字节切片
	salt := make([]byte, length)

	// 使用 crypto/rand 包填充字节切片
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// 将字节切片转换为十六进制字符串
	saltHex := hex.EncodeToString(salt)

	return saltHex, nil
}
