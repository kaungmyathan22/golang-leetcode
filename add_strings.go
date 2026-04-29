package main

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := []byte{}
	for i >= 0 || j >= 0 || carry > 0 {
		var d1, d2 int
		if i >= 0 {
			d1 = int(num1[i] - '0')
		}
		if j >= 0 {
			d2 = int(num2[j] - '0')
		}
		sum := d1 + d2 + carry
		result = append(result, byte(sum%10)+'0')
		carry = sum / 10
		i--
		j--
	}
	// reverse result
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}
	return string(result)
}
