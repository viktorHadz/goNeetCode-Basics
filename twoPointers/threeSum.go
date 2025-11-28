package twopointers

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := range nums {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && a == nums[i-1] {
			continue
		}
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			tripSum := a + nums[lo] + nums[hi]
			if tripSum > 0 {
				hi--
			} else if tripSum < 0 {
				lo++
			} else {
				res = append(res, []int{a, nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			}
		}
	}
	return res
}

// fmt.Printf("\n--- \nnums[i]: %v, nums[lo]: %v, nums[hi]: %v || tripSlice: %v\n---\n\n", nums[i], nums[lo], nums[hi], tripSlice)
