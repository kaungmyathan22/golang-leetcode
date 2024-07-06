package main

func passThePillow(n int, time int) int {
	startingPerson := 1
	add_value := 1
	for t := 0; t < time; t++ {
		startingPerson += add_value
		if startingPerson >= n {
			add_value = -1
		}
		if startingPerson <= 1 {
			add_value = 1
		}
	}
	return startingPerson
}

// func main() {
// 	fmt.Println(passThePillow(4, 5))
// 	fmt.Println(passThePillow(3, 2))
// }
