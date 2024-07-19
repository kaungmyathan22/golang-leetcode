package main

func IsIsomorphic(s string, t string) bool {
	x := make(map[byte]byte)
	y := make(map[byte]byte)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(t); i++ {
		if v1, ok := x[t[i]]; ok && v1 != s[i] {
			return false
		}
		x[t[i]] = s[i]

		if v2, ok := y[s[i]]; ok && v2 != t[i] {
			return false
		}
		y[s[i]] = t[i]
	}
	return true
}
