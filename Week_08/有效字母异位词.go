package Week_08

//计数排序变形
//切片计数
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] += 1
		m[t[i]-'a'] -= 1
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
}

//哈希表计数
func isAnagram0(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := range s {
		m[s[i]] += 1
		m[t[i]] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
