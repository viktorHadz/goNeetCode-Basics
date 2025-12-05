package stack

import "fmt"

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

	open := []string{}
	closed := []string{}

	allowed := map[string]bool{
		"(": true, "{": true, "[": true,
		")": true, "}": true, "]": true,
	}
	allowedOpen := map[int]string{
		0: "(", 1: "{", 2: "[",
	}
	allowedClosed := map[int]string{
		0: ")", 1: "}", 2: "}",
	}

	for _, v := range s {
		val := string(v)
		fmt.Printf("Val %v, \n", val)

		if !allowed[val] {
			return false
		}

		if allowedOpen[0] == val || allowedOpen[1] == val || allowedOpen[2] == val {
			open = append(open, val)
		}
		if allowedClosed[0] == val || allowedClosed[1] == val || allowedClosed[2] == val {
			closed = append(closed, val)
		}

	}
	fmt.Println("Open: ", open)
	fmt.Println("Closed: ", closed)
	return true
}
