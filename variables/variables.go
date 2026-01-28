package main

import "fmt"

func main() {
	var name string = "Alice"
	fmt.Println(name)

	var age int = 30
	fmt.Println(age)

	var isStudent bool = false
	fmt.Println(isStudent)

	var firstName, lastName string = "Vishan", "Chathuranga" // Multiple variable declaration
	fmt.Println("Full Name:", firstName, lastName)

	newUser := "vishan Chathuranga ....."  // Type inference. we can define variable without specifying type
	fmt.Println(newUser)

	var firsNumber int = 30
	var secondNumber int = 20
	var sum int = firsNumber + secondNumber
	fmt.Println("Sum:", sum)

	var userScore float32 = 95.5      // Example of float32 variable
	var userBalance float64 = 1500.75 // Example of float64 variable

	fmt.Println("User Score:", userScore)
	fmt.Println("User Balance:", userBalance)

	//print type
	fmt.Printf("Type of userScore: %T\n", userScore)
	fmt.Printf("Type of userBalance: %T\n", userBalance)
	fmt.Printf("Type of name: %T\n", name)


	// Constant variable
	const pi float64 = 3.14159 // Constant variable declaration not allowed to change value
	fmt.Println("Value of Pi:", pi)



}
