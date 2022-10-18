package main

func main() {

}

func maxProfit(prices []int) int {
	sum, maxProfit := 0, 0
	for index, price := range prices {
		if index+1 < len(prices) {
			sum += prices[index+1] - price
		}
		if sum < 0 {
			sum = 0
		}
		maxProfit = max(sum, maxProfit)
	}
	return maxProfit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
