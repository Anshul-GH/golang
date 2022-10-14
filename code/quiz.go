package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	// var name string = "Tim"
	// var age int = 32
	// var height float64 = 5.67
	// name := "Tim"
	// age := 32
	// height := 5.67

	// fmt.Printf("Hello %v, how are you today?", name)
	// fmt.Printf("Hello %v, you are %v years old.", name, age)
	// fmt.Println(height)

	// var name string
	// var age int

	// fmt.Scan(&name)
	// fmt.Scan(&age)

	// fmt.Println(name)
	// fmt.Println(age)

	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!", name)

	var age uint
	fmt.Printf("\nEnter your age: ")
	fmt.Scan(&age)

	var eligible bool

	eligible = age >= 10

	if eligible {
		fmt.Println("Congrats! You are eligible to play the game.")
	} else {
		fmt.Println("Sorry .. You are NOT eligible to play the game.")
	}
}
