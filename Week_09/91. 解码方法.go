package Week_09

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] < '7' {
			dp[i] = dp[i] + dp[i-2]
		}
	}
	return dp[n]
}

//dp中i为计算的字符数
//s中i为计算第i位字符的索引
//当s中当前位前一位不为'0'时，i个字符的解码方法数与i-1个数的解码方法数一致；
//当s中当前位前2位为'1'或'2'时，且i-1位的值为0-6中的一个值，当前数量字符的解码方法数，为i个字符数的解码方法数+i-2个字符数的解码方法数。
//最终结果取字符串的长度值。
