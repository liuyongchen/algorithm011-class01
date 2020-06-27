package Week_01

func MaxArea(height []int) int {
	maxA := 0
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		maxA = max(maxA, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
