package main

func isHappy(n int) bool {
	m := make(map[int]int)
	for true {
		n = NumSquareSum(n)
		if n == 1 {
			return true
		}
		if _, exist := m[n]; exist {
			return false
		}
		m[n] = 1
	}

	return false
}

func NumSquareSum(n int) int {
	sum := 0
	for n != 0 {
		a := n % 10
		n = n / 10
		sum += a * a
	}
	return sum
}
