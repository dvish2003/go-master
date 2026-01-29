package main

import "fmt"

func main(){
	//arrays
	var numbers [5]int = [5]int{10,20,30,40,50} //array declaration and initialization
	fmt.Println("Array:", numbers)

	//accessing array elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])


	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Colors Array:", colors)

	//length of array
	fmt.Println("Length of numbers array:", len(numbers))
	fmt.Println("Length of colors array:", len(colors))
}


