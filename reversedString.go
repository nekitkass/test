package main

func ReverseString(str string) (reversed string) {
	for _, char := range str {
		reversed = string(char) + reversed
	}
	return
}
