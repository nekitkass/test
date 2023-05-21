package main

func SmallestIntegerFinder(numbers []int) int {
	result := numbers[0]
	for _, value := range numbers {
		if value < result {
			result = value
		}
	}
	return result
}
