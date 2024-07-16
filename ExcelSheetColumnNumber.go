package main

func titleToNumber(columnTitle string) int {
	var result = 0

	for _, char := range columnTitle {
		intValue := int(char)
		result *= 26
		result += intValue - 'A' + 1
	}
	return result
}
