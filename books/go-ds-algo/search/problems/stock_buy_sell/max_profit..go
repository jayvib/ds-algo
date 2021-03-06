package stock_buy_sell

import "math"

func MaxProfit(price []int, start, end int) int {

	// Can't buy stock from the past price
	// 23 40 33
	//    e<-s
	if end <= start {
		return 0
	}

	var profit int

	// []int{100, 180, 260, 310, 40, 535, 695}
	// The day at which the stock must be bought
	for i := start; i < end; i++ {

		// The day at which the stock must be sold
		for j := i + 1; j <= end; j++ {

			// If buying the stock at ith day and
			// selling it at jth day is profitable
			if price[j] > price[i] {
				currentProfit := price[j] - price[i] +
					MaxProfit(price, start, i-1) + // Calculate the previous profit from the starting day up to this day
					MaxProfit(price, j+1, end) // Calculate the profit from this day and onward
				profit = int(math.Max(float64(profit), float64(currentProfit)))
			}
		}
	}

	return profit
}

func maxProfit(price []int, start, end int) int {
	// Base case: Can't calculate the price of the
	// stock that haven't been bought
	//if end <= start {
	//	return 0
	//}

	var profit int
	//for i := start; i < end; i++ {

	//}
	return profit
}
