package main

import "fmt"

// function to greet the user
func greet() {
	fmt.Println("Hi, my name is Alim Ikegami! Nice to meet you!")
}

func askQuestion() {
	fmt.Println("How are you today?")
}

func main() {
	fmt.Println("Hello world")
	greet()
	askQuestion()
}
