package arraysandhashing

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		req := target - num
		if j, ok := m[req]; ok {
			// j < i guaranteed by construction
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// func TwoSum(nums []int, target int) []int {
// 	mNums := make(map[int]int)
// 	for idx, num := range nums {
// 		mNums[num] = idx
// 	}

// 	for idx, num := range nums {
// 		req := target - num

// 		_, exists := mNums[req]
// 		if exists {
// 			if idx == mNums[req] {
// 				continue
// 			}
// 			if mNums[req] > idx {
// 				return []int{idx, mNums[req]}
// 			} else {
// 				return []int{mNums[req], idx}
// 			}
// 		}
// 	}
// 	return []int{0, 0}
// }
