package binarysearch

import "fmt"

// Neetcode binary search easy prob 1

func Search(nums []int, target int) int {
	s := 0
	e := len(nums) - 1

	fmt.Printf("\ntarget: %v\nstart: %v end: %v\nnums: %v\n", target, s, e, nums)
	for s <= e {
		mid := (s + e) / 2

		fmt.Printf("\ntarget: %v\nstart: %v end: %v\nmid: %v | nums[mid]: %v\nnums: %v\n", target, s, e, mid, nums[mid], nums)

		if target < nums[mid] {
			e = mid - 1
		}
		if target > nums[mid] {
			s = mid + 1
		}
		if nums[mid] == target {
			return mid
		}
	}
	fmt.Println("Target not found: -1")
	return -1
}
