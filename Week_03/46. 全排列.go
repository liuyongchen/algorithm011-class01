package Week_03

//简洁写法
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{nums}
	}
	res := [][]int{}

	for i, num := range nums {
		temp := make([]int, len(nums))
		copy(temp, nums[:i])
		copy(temp[i:], nums[i+1:])
		sub := permute(temp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}

//回溯法经典写法 思路很清晰.
func permute1(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	path := []int{}
	visited := make([]bool, len(nums))
	dfshelp(nums, len(nums), 0, path, visited, &res)
	return res
}

func dfshelp(nums []int, lenth int, depth int, path []int, visited []bool, res *[][]int) {
	if depth == len(nums) {
		*res = append(*res, path)
		return
	}
	for k, v := range nums {
		if visited[k] {
			continue
		}
		//path 需要每次重新定义
		path := append(path, v)
		visited[k] = true
		dfshelp(nums, lenth, depth+1, path, visited, res)
		path = path[:len(path)-1]
		visited[k] = false
	}
}
