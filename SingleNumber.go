package main

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			delete(m, nums[i])
		} else {
			m[nums[i]] = nums[i]
		}
	}

	for key, _ := range m {
		return key
	}
	return 0
}
