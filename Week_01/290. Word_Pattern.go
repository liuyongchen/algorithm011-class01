package Week_01

import "strings"

func wordPattern(pattern string, str string) bool {
	str_list := strings.Split(str, " ")
	if len(pattern) != len(str_list) {
		return false
	}
	patternM, strM := make(map[byte]int), make(map[string]int)
	for i := range pattern {
		if _, ok := patternM[pattern[i]]; !ok {
			patternM[pattern[i]] = i
		}
		if _, ok := strM[str_list[i]]; !ok {
			strM[str_list[i]] = i
		}
		if patternM[pattern[i]] != strM[str_list[i]] {
			return false
		}
	}
	return true
}
