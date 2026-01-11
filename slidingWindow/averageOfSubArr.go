package slidingwindow

func AverageOfSubarrays(arr []int, k int) []float64 {
	if k <= 0 || k >= len(arr) {
		panic("Fix function inputs. Wrong size of param k.")
	}
	answer := []float64{}

	var tempAvg float64
	for i := range k {
		tempAvg += float64(arr[i])
	}
	tempAvg = float64(tempAvg) / float64(k)
	answer = append(answer, float64(tempAvg))
	// Build sliding window
	for right := k; right < len(arr); right++ {
		tempAvg += float64(arr[right])
		tempAvg -= float64(arr[right-k])
		avg := tempAvg / float64(k)
		answer = append(answer, float64(avg))

	}
	// Calculate average & add to averages[]
	return answer
}
