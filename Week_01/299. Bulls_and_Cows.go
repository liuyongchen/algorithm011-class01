package Week_01

import (
	"fmt"
	"strconv"
)

func GetHint0(secret string, guess string) string {
	bull := 0
	cow := 0
	m := make(map[int32]int)
	for idx, vel := range secret {

		if int32(guess[idx]) == vel {
			bull++
		}
		m[vel]++
	}
	for _, gvel := range guess {
		if m[gvel] != 0 {
			cow++
			m[gvel]--
		}
	}
	res := fmt.Sprintf("%dA%dB", bull, cow-bull)
	return res
}

func GetHint(secret string, guess string) string {
	m := make([]int, 128)
	a, b := 0, 0
	l := len(secret)
	for i := 0; i < l; i++ {
		m[secret[i]]++
	}
	for i := 0; i < l; i++ {
		n := m[guess[i]]
		if secret[i] == guess[i] {
			if n == 0 {
				b--
			} else {
				m[guess[i]]--
			}
			a++
		} else {
			if n > 0 {
				b++
				m[guess[i]]--
			}
		}

	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
