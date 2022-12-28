package src

import "strings"

func IsIn(phrase string, str string) bool {

	p := strings.ToLower(phrase)
	s := strings.ToLower(str)

	lPhr := len(phrase)
	lStr := len(str)

	if lStr < lPhr {
		return false
	}

	for i := 0; i < lStr; i++ {
		if p[0] == s[i] {
			for j := i; j < lStr; j++ {
				if s[j] == p[j-i] {
					if j-i == lPhr-1 {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}
