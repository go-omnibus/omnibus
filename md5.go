package omnibus

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
