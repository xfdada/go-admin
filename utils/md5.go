package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// HashAndSalt 加密密码
func HashAndSalt(pwd string) string {
	pwds := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwds, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords 验证密码
func ComparePasswords(hashedPwd, Pwd string) bool {
	byteHash := []byte(hashedPwd)
	Pwds := []byte(Pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, Pwds)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
