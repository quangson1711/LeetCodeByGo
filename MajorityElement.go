package main

func majorityElement(nums []int) int {
	temps := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		value, ok := temps[nums[i]]
		if ok {
			temps[nums[i]] = value + 1
			if temps[nums[i]] > len(nums)/2 {
				return nums[i]
			}
		} else {
			temps[nums[i]] = 1
			if temps[nums[i]] > len(nums)/2 {
				return nums[i]
			}
		}
	}
	return -1
}
