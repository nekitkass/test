package main

func FindOdd(seq []int) int {
	matches := make(map[int]int)
	for _, number := range seq {
		if value, ok := matches[number]; ok {
			matches[number] = value + 1
		} else {
			matches[number] = 1
		}
	}
	for key, value := range matches {
		if value%2 != 0 {
			return key
		}
	}
	return 0
}
