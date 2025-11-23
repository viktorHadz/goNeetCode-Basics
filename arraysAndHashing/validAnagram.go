package arraysandhashing

import (
	"fmt"
	"maps"
)

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	setS := make(map[rune]int)
	setT := make(map[rune]int)

	for _, r := range s {
		setS[r]++
	}
	for _, r := range t {
		setT[r]++
	}
	fmt.Println(s)
	fmt.Println(setS, "\n---")
	fmt.Println(t)
	fmt.Println(setT, "\n---")
	return maps.Equal(setS, setT)
}
