package main

import slidingwindow "github.com/viktorHadzh/goNeetCode-Basics/slidingWindow"

// lowercase var => private
// uppercase var => public
func main() {
	// nonNeet.RomanToInt("XLVIII")
	slidingwindow.MaxProfit([]int{10, 1, 5, 6, 7, 1})
	// binarysearch.SearchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10)
	// binarysearch.Search([]int{-1, 0, 2, 4, 6, 8}, 3)

	// --- Reversed order. New code goes above
	// fmt.Println(arrHash.HasDuplicate([]int{1, 2, 3, 3}))
	// fmt.Println(arrHash.IsAnagram("bbcca", "ccbca"))
	// fmt.Println(arrHash.TwoSum([]int{5, 5}, 10))
	// fmt.Println(arrHash.GroupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
	// fmt.Println(twopointers.AlphaNumJoin("Was it a car or a cat I saw?"))
	// fmt.Println(twopointers.TwoSum([]int{1, 2, 3, 4}, 3))
	// fmt.Println(twopointers.IsPalindrome2("tab a cat"))
	// fmt.Println(twopointers.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(twopointers.ThreeSum([]int{0, 0, 0, 0}))
	// fmt.Println(stack.IsValidParenthesis("(){}}{"))

	// Using my custom stack DS
	// stack := stack.NewStack[string]()
	// stack.Push("1. One")
	// stack.Push("2.Two")
	// stack.Push("3. Three")
	// fmt.Println(stack)
	// v, ok := stack.Pop()
	// if !ok {
	// 	return
	// } else {
	// 	fmt.Println("Removing: ", v)
	// }
	// fmt.Println("Stack Size:", stack.GetStackSize())

}
