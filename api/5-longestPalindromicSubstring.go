package api

func longestPalindrome(s string) string {

	var result string

	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s), len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && (len(result) == 0 || j-i+1 > len(result)) {
				result = s[i : j+1]
			}
		}
	}

	return result
}
