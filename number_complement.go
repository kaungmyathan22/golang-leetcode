package main

import (
	"strconv"
)

func findComplement(num int) int {
	// convert num to binary string
	binaryStr := (strconv.FormatInt(int64(num), 2))
	// loop through it and reverse it
	revertedStr := ""
	for _, v := range binaryStr {
		if string(v) == "0" {
			revertedStr += "1"
		} else {
			revertedStr += "0"
		}
	}
	// convert reverted str to num
	revertedNum, _ := strconv.ParseInt(revertedStr, 2, 64)
	return int(revertedNum)
}
