package conver

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func StringToMD5(str string) string {
	lowStr := strings.ToLower(str)
	hash := md5.Sum([]byte(lowStr))
	return hex.EncodeToString(hash[:])
}
