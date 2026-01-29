package main

import "fmt"


//function increment use pointer

func increment(number *int) {
   *number ++
}

//function to set Zero
func setZero(number *int){
	*number = 0
	}
	
func main(){

	 num := 10
	fmt.Println("Before increment:", num)
	increment(&num)
	fmt.Println("After increment:", num)
	
	setZero(&num)
	fmt.Println("After setZero:", num)
}
