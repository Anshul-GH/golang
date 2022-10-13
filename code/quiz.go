package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	// var name string = "Tim"
	// var age int = 32
	// var height float64 = 5.67
	name := "Tim"
	age := 32
	height := 5.67

	fmt.Printf("Hello %v, how are you today?", name)
	fmt.Printf("Hello %v, you are %v years old.", name, age)
	fmt.Println(height)
}
