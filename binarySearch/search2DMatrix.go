package binarysearch

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		fmt.Println("Row", row)
		lastInRow := row[len(row)-1]
		fmt.Println("Last num: ", lastInRow)

		if lastInRow >= target {
			fmt.Printf("-->Target found in: %v\nLast in row %v, Target: %v\n", row, lastInRow, target)
			s := 0
			e := len(row) - 1
			fmt.Printf("start: %v end: %v\n", s, e)
			for s <= e {
				mid := (s + e) / 2
				if row[mid] > target {
					fmt.Println("Row mid", row[mid])
					e = mid - 1
				}
				if row[mid] < target {
					fmt.Println("Row mid", row[mid])
					s = mid + 1
				}
				if row[mid] == target {
					fmt.Println("---Target-->", row[mid])
					return true
				}
			}
			return false
		} else {
			continue
		}

	}
	return false
}
