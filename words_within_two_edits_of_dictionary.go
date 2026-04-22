package main

func twoEditWords(queries []string, dictionary []string) []string {
	result := make([]string,0)
	for q := range queries {
		for d := range dictionary {
			if isWithinTwoEdits(queries[q], dictionary[d]) {
				result = append(result, queries[q])
				break
			}
		}
	}
	return result
}

func isWithinTwoEdits(query string, dictionary string) bool {
	diff := 0
	for i := 0; i < len(query); i++ {
		if query[i] != dictionary[i] {
			diff++
		}
	}
	return diff <= 2
}