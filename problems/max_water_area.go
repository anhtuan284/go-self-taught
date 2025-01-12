package problems

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// Calculate the width and the height of the container
		width := right - left
		h := min(height[left], height[right])

		// Calculate the area and update the maximum area
		area := width * h
		if area > maxWater {
			maxWater = area
		}

		// Move the pointer of the shorter line inward
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
