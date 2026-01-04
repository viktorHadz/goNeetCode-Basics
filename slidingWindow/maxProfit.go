package slidingwindow

import "fmt"

// Prices: {10, 1, 5, 10, 6, 7, 1}
// l = 2
// r = 3
// 1. prices[l] = 10 prices[r] = 1
// 2. prices[l] = ? prices[r] = 5
func MaxProfit(prices []int) int {
	l, r := 0, 1 // l==buy, r==sell 
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit>maxProfit {
				maxProfit = profit
			}
		} else {
			l = r
		}
		fmt.Printf("buy: %v sell: %v, maxProfit: %v\n", prices[l],prices[r], maxProfit)
		r+=1
	}
	return maxProfit
}
