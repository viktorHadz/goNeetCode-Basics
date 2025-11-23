package arraysandhashing

import (
	"fmt"
)

func HasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, n := range nums {
		_, exists := seen[n]
		if exists {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}

func PrintMap() {
	myMap := make(map[string]int)
	myMap["pens"] = 2
	myMap["book"] = 4
	myMap["notebooks"] = 5

	val, ok := myMap["22"]

	fmt.Println("ok: ", ok)
	fmt.Println("val: ", val)
	fmt.Println(myMap)

}
