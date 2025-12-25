package slidingwindow

import "fmt"

// Example {10, 1, 5, 6, 7, 1}
func MaxProfit(prices []int) int {
	answer := 0
	low := 0
	high := 0

	for i := 0; i+1 < len(prices); i++ {
		p1 := prices[i]
		p2 := prices[i+1]

		if p1 > p2 {
			low = p2
			fmt.Println("Low new val: ", low)
		}
		if p1 < p2 {
			high = p2
			fmt.Println("High new val: ", high)
		}

		if high > low {
			answer = high - low
		}

		fmt.Printf("Start: %v End: %v \n", p1, p2)
	}

	fmt.Println("Answer: ", answer)
	return answer
}
