package main

import (
	"fmt"

	arrHash "github.com/viktorHadzh/goNeetCode-Basics/arraysAndHashing"
)

// lowercase var => private
// uppercase var => public
func main() {
	// fmt.Println(arrHash.HasDuplicate([]int{1, 2, 3, 3}))
	fmt.Println(arrHash.IsAnagram("bbcca", "ccbca"))

}
