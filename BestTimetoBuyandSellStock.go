package main

import "log"

func maxProfit2(prices []int) int {
	log.Println("start")
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				tempProfit := prices[j] - prices[i]
				if tempProfit > profit {
					profit = tempProfit
				}
			}
		}
	}
	log.Println("end")
	return profit
}

func maxProfit(prices []int) int {
	log.Println("start")
	profit := 0
	tempMin := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < tempMin {
			tempMin = prices[i]
		} else {
			tempProfit := prices[i] - tempMin
			if tempProfit > profit {
				profit = tempProfit
			}
		}
	}
	log.Println("end")
	return profit
}
