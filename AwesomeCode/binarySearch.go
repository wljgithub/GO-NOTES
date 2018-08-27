package main

import (
	"fmt"
	"sort"
)

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	fmt.Println('y')
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != ""
	})
	fmt.Printf("Your number is %d.\n", answer)
}
func main() {

	GuessingGame()
}
