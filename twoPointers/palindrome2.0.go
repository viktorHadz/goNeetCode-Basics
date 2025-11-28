package twopointers

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome2(s string) bool {
	var b strings.Builder
	for _, v := range s {
		if unicode.IsSpace(v) || unicode.IsPunct(v) {
			continue
		}
		b.WriteRune(unicode.ToLower(v))
		fmt.Println(string(v))
	}
	str := b.String()
	fmt.Println(str)

	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
