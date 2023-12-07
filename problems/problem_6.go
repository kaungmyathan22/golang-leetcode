package problems

import (
	"strconv"
	"strings"
)

func LargestOddNumber(num string) string {
	str := strings.Split(num, "")
	for i := len(str) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(str[i])
		if val%2 == 0 {
			continue
		} else {
			if i == 0 {
				return str[0]
			}
			return strings.Join(str[0:i+1], "")
		}
	}
	return ""
}
