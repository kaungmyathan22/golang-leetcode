func reverseVowels(s string) string {
	vowelSet := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	left := 0
	right := len(s)
	b := []byte(s)
	var err error
	for left < right {
		for left < right && !vowelSet[s[left]] {
			left++
		}
		for left < right && !vowelSet[s[right]] {
			right--
		}
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}
