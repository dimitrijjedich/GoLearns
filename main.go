package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Printf("Hello, %s! You are %d years old!\n", name, age)
	if age < 18 {
		var missing int = 18 - age
		var years string
		if missing == 1 {
			years = "year"
		} else {
			years = "years"
		}
		fmt.Printf("You are underage and it takes %d %s until you are 18", missing, years)
	} else {
		fmt.Printf("You are not underage!")
	}
}
