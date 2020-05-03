package lib

import (
	"crypto/sha1"
	"strconv"
)

func StringToUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}

func UintToString(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

// sha1S SHA1散列
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}
