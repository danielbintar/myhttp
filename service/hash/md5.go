package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(link string) string {
	hashed := md5.Sum([]byte(link))
	return hex.EncodeToString(hashed[:])
}
