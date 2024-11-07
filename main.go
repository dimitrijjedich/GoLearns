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
	fmt.Printf("Hello, %s! You are %d years old!", name, age)
	if age < 18 {
		fmt.Printf("You are underage and it takes %d years until you are 18", 18-age)
	} else {
		fmt.Printf("You are not underage!")
	}
}
