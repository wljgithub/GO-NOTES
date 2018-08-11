package main

import (
	"fmt"
)

var (
	firstName, lastName string
	hobby               string
	age                 string
)

func main() {
	fmt.Println("enter your first name and last name separately please!")
	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)
	fmt.Scanf("%s,%s", &hobby, &age)

	fmt.Println("full name :", firstName, lastName)
	fmt.Println("and hobby is ", hobby)
	fmt.Println(age, "years old")
}
