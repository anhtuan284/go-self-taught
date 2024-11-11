package problems

func CheckRegExMatch(s string, p string) bool {
	// Initialize the DP table with dimensions (len(s)+1) x (len(p)+1)
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true // Base case: empty pattern matches empty string

	// Handle cases where '*' appears in the pattern and can match empty strings
	for j := 2; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the DP table
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1] // Direct match or match by '.'
			} else if p[j-1] == '*' {
				// Assume '*' represents zero occurrences of the preceding element
				dp[i][j] = dp[i][j-2]
				// If preceding character matches s[i-1] or is '.', match one or more occurrences
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	// Final result: does s match p entirely?
	return dp[len(s)][len(p)]
}
