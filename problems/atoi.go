package problems

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)
	sign := 1
	result := 0

	// Ignore leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	// Handle sign if it exists
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert number and avoid overflow
	for i < n && unicode.IsDigit(rune(s[i])) {
		digit := int(s[i] - '0')

		// Check for overflow before updating result
		if result > (math.MaxInt32-digit)/10 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
