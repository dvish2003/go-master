package main

import "fmt"

func main() {
	  numbers := []int{10, 20, 30, 40, 50,32,32,33,43,45,46,5} //slice declaration and initialization
	  fmt.Println("Original Slice:", numbers)

	  numbers = append(numbers, 90,2343,34,324324,324,324,2)
	  fmt.Println("After appending", numbers)

	  //remove
	  numbers2 := []int{1,2,3,4,5,6,7,8,9,10}
	  fmt.Println("Before removing", numbers2)
	   numbers2 = append(numbers2[:3], numbers2[4:]...) //remove 4th element
	   fmt.Println(numbers2)//remove 4th element
	  
}
