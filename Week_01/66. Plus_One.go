package Week_01

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	res := make([]int, len(digits)+1)
	if digits[0] == 0 {
		res[0] = 1
		for i, v := range digits[:] {
			res[i+1] = v
		}
		return res
	}
	return digits
}
