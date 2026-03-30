package main

import "strings"

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	result := []string{}

	for _, word := range words {
		lowerWord := strings.ToLower(word)

		// Figure out which row the first letter belongs to
		firstChar := string(lowerWord[0])
		rowIndex := -1

		for i, row := range rows {
			if strings.Contains(row, firstChar) {
				rowIndex = i
				break
			}
		}

		// Check if all letters are in the same row
		valid := true
		for _, c := range lowerWord {
			if !strings.Contains(rows[rowIndex], string(c)) {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, word)
		}
	}

	return result
}
