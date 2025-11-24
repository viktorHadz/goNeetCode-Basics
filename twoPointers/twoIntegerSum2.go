package twopointers

func TwoSum(numbers []int, target int) []int {
	L := 0
	R := len(numbers) - 1

	for L < R {
		sum := numbers[L] + numbers[R]
		// Return if n[L] + n[R] == target
		if sum == target {
			return []int{L + 1, R + 1}
		}

		// Move pointers if != target
		if sum < target {
			L++
		} else {
			R--
		}
	}

	return nil
}
