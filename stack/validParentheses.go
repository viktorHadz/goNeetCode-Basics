package stack

import (
	"fmt"
)

// 1. Make 2 stacks => 1. Open 2. Closed
// 2. Define accepted symbols
// 1. Open ["(", "{", "["]
// 2. Closed [")", "}", "]"]
// 3. Loop over sting
// 4. Conditions
// 1. if string !== "(" "{" "[" or "]" "}" ")" return false
// 2. If string == "(" "{" "[" push in 1. Open
// 3. If string == ")" "}" "]" push in 2. Closed
// 5. Record position of each (map?)
// Add string 1. Open [( : 0 { : 1 [ : 2]  2. Closed [ ) : 0 } : 1 ] : 2]

// Add open current ==> then check if closed current position !not empty
// if open 0 != "empty" and closed 0 != "empty"
// Compare Open 0 to Closed 0
// If Open !== Closed reeturn false
// Else continue

//------
/* if closed == true {
	check if open at pos[closed]
	return
}
*/

func IsValidParenthesis(s string) bool {
	if len(s)%2 != 0 {
		fmt.Println("len(s)%2 != 0 FALSE")
		return false
	}

	pairs := map[byte]byte{
		'(': ')', '{': '}', '[': ']',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			stack = stack[:ch]
		}
	}
}
