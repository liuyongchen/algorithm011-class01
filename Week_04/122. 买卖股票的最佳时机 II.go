package Week_04

func maxProfit(prices []int) int {
	var res int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

// 后一天的价格如果比前一天高，就计算两天的差价，
// 直到遇到后一天比前一天的价格低，停止计算。
// 最终得到了，股票价格一直上涨的收益和。
