package problems

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	reversed := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && pop < -8) {
			return 0
		}

		reversed = reversed*10 + pop
	}
	return reversed
}

func main() {
	a := 1231412
	fmt.Print(reverse(a) == 2141321)
}
