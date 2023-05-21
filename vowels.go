package main

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for char := range str {
		for vowel := range vowels {
			if string(str[char]) == vowels[vowel] {
				count++
			}
		}
	}
	return count
}
