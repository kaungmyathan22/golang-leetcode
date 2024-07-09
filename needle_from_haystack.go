package main

func NeedleFromHaystack(haystack string, needle string) int {
OuterLoop:
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for ii := 0; ii < len(needle); ii++ {
				if i+ii >= len(haystack) || haystack[i+ii] != needle[ii] {
					continue OuterLoop
				}
			}
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(NeedleFromHaystack("sadbutsad", "sad"))
// 	fmt.Println(NeedleFromHaystack("leetcode", "leeto"))
// }
