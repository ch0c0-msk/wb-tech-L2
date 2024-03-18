package dev02

import (
	"errors"
	"unicode"
)

// UnpackString unzips the RLE string
func UnpackString(common string) (string, error) {
	res := make([]rune, 0)
	runeCommon := []rune(common)
	for i := 0; i < len(runeCommon); i++ {
		if unicode.IsDigit(runeCommon[i]) {
			if i == 0 {
				return string(res), errors.New("string cannot start with a digit")
			}
			repeatCount := int(runeCommon[i] - '0')
			for ; i < len(runeCommon)-1 && unicode.IsDigit(runeCommon[i+1]); i++ {
				repeatCount = repeatCount*10 + int(runeCommon[i]-'0')
			}
			repeatCount--
			buf := make([]rune, repeatCount)
			for j := range buf {
				buf[j] = res[len(res)-1]
			}
			res = append(res, buf...)
		} else if runeCommon[i] == '\\' {
			res = append(res, runeCommon[i+1])
			i++
		} else {
			res = append(res, runeCommon[i])
		}
	}
	return string(res), nil
}
