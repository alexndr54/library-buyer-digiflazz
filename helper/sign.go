package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5Hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
