package Week_04

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

// 逐个遍历能满足孩子胃口的饼干，
// 只要遍历到刚好能满足孩子胃口的饼干，i++,继续满足下一个孩子，
// 当所有饼干不能满足下一个孩子的时候，后续更大胃口的孩子也不能继续满足，返回i的值
