package slidingwindow

import "fmt"

// Example {10, 1, 5, 6, 7, 1}
func MaxProfit(prices []int) int {
	p1 := 0
	p2 := len(prices) - 1
	answer := 0
	for p1 < p2 {
		if prices[p1] > prices[p2] {
			p1++
		} else {
			p2--
		}
	}
	fmt.Println("Answer: ", answer)
	return answer
}
