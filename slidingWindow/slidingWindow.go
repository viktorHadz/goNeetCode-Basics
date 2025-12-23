package slidingwindow

import (
	"fmt"
	"slices"
)

func MaxProfit(prices []int) int {
	// given an array of ints
	// where each int is the price of the coin on that day
	// do: have to buy/sell so that i make profit

	// ok so the idea is buy low sell high
	// Make 2 pointers one for start the other for finish
	// If end is
	fmt.Println("Prices: ", prices)
	slices.Sort(prices)
	fmt.Println("Sorted: ", prices)
	s := 0
	e := len(prices) - 1
	for i:=0; i <= len(prices); i++ {
		
	}
	return 1
}
