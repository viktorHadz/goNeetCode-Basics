package slidingwindow

func MaxSumOfK(arr []int, k int) int {
	if k <= 0 || k > len(arr) {
		panic("Param k <= 0 || k >= slice (slc []int). Reduce inputs in main.")
	}

	tempSum := 0
	for i := range k {
		tempSum += arr[i]
	}
	maxSum := tempSum
	for right := k; right < len(arr); right++ {
		tempSum += arr[right]
		tempSum -= arr[right-k]

		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return  maxSum
}
