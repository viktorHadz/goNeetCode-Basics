package stack

import (
	"fmt"
)

func IsValidParenthesis(s string) bool {
	if len(s)%2 != 0 {
		fmt.Println("len(s)%2 != 0 FALSE")
		return false
	}

	pairs := map[byte]byte{
		')': '(', '}': '{', ']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
			fmt.Println("Pushing: ", string(stack))
		case ')', '}', ']':

			fmt.Println("Pairs[ch]: ", string(pairs[ch]))
			// If pairs[ch] "(" is at the top of the stack
			if len(stack) != 0 && pairs[ch] == stack[len(stack)-1] {
				lastEl := len(stack) - 1 // idx last el
				// pop from stack
				stack = stack[:lastEl]
			} else {
				return false
			}
		}
	}
	fmt.Println("Stack: ", string(stack))
	// If stack is == 0 return true
	if len(stack) != 0 {
		return false
	} else {
		return true
	}

}
