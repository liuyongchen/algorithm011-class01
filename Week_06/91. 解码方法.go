package Week_06

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
