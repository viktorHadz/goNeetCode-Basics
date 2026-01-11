package nonneet

func RomanToInt(s string) int {
	r := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	answer := 0
	for i := 0; i < len(s); i++ {
		v := r[s[i]]

		if i+1 < len(s) && v < r[s[i+1]] {
			answer -= v
		} else {
			answer += v
		}
	}
	return answer
}
