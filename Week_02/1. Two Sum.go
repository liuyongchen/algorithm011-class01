package Week_02

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		need := target - val
		if idxN, ok := m[need]; ok {
			if idx != idxN {
				return []int{idx, idxN}
			}
		}
		m[val] = idx
	}
	return []int{}
}

//枚举 a,b --> a + b == target -- >
//for each a: check target= a exits in nums
