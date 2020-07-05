package Week_02

import "sort"

//哈希表
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		if m[v] > 0 {
			nums2[k] = v
			m[v]--
			k++
		}
	}
	return nums2[:k]
}

func intersect1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]] += 1
	}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		if m[nums1[i]] > 0 {
			res = append(res, nums1[1])
			m[nums1[i]]--
		}
	}
	return res
}

//双指针
func intersect2(nums1 []int, nums2 []int) []int {
	i, j, res := 0, 0, []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
			continue
		}
		if nums2[j] > nums1[i] {
			i++
			continue
		}
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}
