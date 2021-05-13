package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(datum string) string {
	hashed := md5.Sum([]byte(datum))
	return hex.EncodeToString(hashed[:])
}
