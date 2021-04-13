package 数组和字符串

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	// dp[i][j] 表示 i ~ j 是否时回文串
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	start := 0
	maxLen := 1
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			//
			if s[i] != s[j] {
				dp[i][j] = 0
			} else {
				if j-i <= 2 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			 // dp[i][j]是回文串
			if dp[i][j] > 0 && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return s[start : start+maxLen]
}
