package main

func Grow(arr []int) int {
	var result = 1
	for index := range arr {
		result = arr[index] * result
	}
	return result
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	counter := 0
	for p0 < p {
		p0 = p0 + int(float64(p0)*percent/100) + aug
		counter++
	}
	return counter
}
