package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

var (
	salt = []byte{225, 229, 205, 212, 6, 107, 168, 7, 60, 90, 113, 26, 220, 30, 107, 117, 17, 132, 252, 243, 77, 73, 154, 56, 208, 216, 74, 198, 24, 1, 73, 73}
	iter = 1024
)

// CryptPasswd 用户密码设计：客户端/前端加密用 MD5 加密，后端再用 pbkdf2 加密
func CryptPasswd(passwd []byte) string {
	dk := pbkdf2.Key(passwd, salt, iter, 32, sha256.New)
	return hex.EncodeToString(dk)
}

func InitSalt(n int) ([]byte, error) {
	salt := make([]byte, n)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
