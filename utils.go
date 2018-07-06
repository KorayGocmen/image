package image

// max finds max among multiple integers
func max(ints ...int) int {
	var max = ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}
	return max
}

// min finds min among multiple integers
func min(ints ...int) int {
	var min = ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}
	return min
}
