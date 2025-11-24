package main

import (
	"fmt"

	twopointers "github.com/viktorHadzh/goNeetCode-Basics/twoPointers"
)

// lowercase var => private
// uppercase var => public
func main() {
	// fmt.Println(arrHash.HasDuplicate([]int{1, 2, 3, 3}))
	// fmt.Println(arrHash.IsAnagram("bbcca", "ccbca"))
	// fmt.Println(arrHash.TwoSum([]int{5, 5}, 10))
	// fmt.Println(arrHash.GroupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
	// fmt.Println(twopointers.AlphaNumJoin("Was it a car or a cat I saw?"))
	fmt.Println(twopointers.TwoSum([]int{1, 2, 3, 4}, 3))

}
