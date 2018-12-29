package day04

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func firstMatch(input, prefix string) int {
	for i := 1; ; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		str := hex.EncodeToString(hash[:])
		if strings.HasPrefix(str, prefix) {
			return i
		}
	}
}
