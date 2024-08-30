package crypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func LoginCrypt(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum(nil))
}
