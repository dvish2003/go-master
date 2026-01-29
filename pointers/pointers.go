package main

import "fmt"

func main(){
	//pointer
	var number int = 42
	var pointerToNumber *int = &number // Pointer variable declaration

	fmt.Println("Value of number:", number)
	fmt.Println("Address of number:", &number)
	fmt.Println("Value of pointerToNumber (address of number):", pointerToNumber)
	fmt.Println("Value at the address stored in pointerToNumber:", *pointerToNumber) // Dereferencing the pointer 

}