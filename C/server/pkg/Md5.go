package pkg

import (
	"crypto/md5"
	"fmt"
)

func PwdMd5(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}
