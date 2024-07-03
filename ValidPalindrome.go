package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = RemoveSpecialSyntax(s)
	a := Reverse(s)
	if a == s {
		return true
	}
	return false
}

func RemoveSpecialSyntax(s string) string {
	// Tạo regex để chỉ lấy số và chữ cái (A-Z, a-z, 0-9)
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)

	// Tìm tất cả các khớp trong chuỗi
	matches := re.FindAllString(s, -1)

	// Kết hợp các kết quả thành một chuỗi duy nhất
	result := ""
	for _, match := range matches {
		result += match
	}
	return result
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
