package twopointers

import (
	"fmt"
	"strings"
	"unicode"
)

func RemoveSpace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			continue
		}

		b.WriteRune(unicode.ToLower(r))
	}

	return b.String()
}

// Lowercases all chracters and removes non alphanumeric characters (" ", "?,!,."...)
func AlphaNumJoin(s string) bool {
	str := RemoveSpace(s)
	L := 0
	R := len(str)
	fmt.Println(str)

	for L < R {
		fmt.Println("Left =>", string(str[L]), "Right =>", string(str[R-1]))
		if str[L] != str[R-1] {
			return false
		}
		L++
		R--

	}

	return true
}

// func IsPalindrome(s string) bool {
// 	strJoin := SanitizeToAlphaNum(s)
// 	if palindrome {
// 		return true

// 	} else {
// 		return false
// 	}
// }
