package main



import "fmt"

func main() {
	 var twoDArrays = [3][3]int {{1,2,3},{4,5,6},{7,8,9}} //2D array declaration and initialization
	 fmt.Println("2D Array:", twoDArrays)

	 //accessing elements
	 fmt.Println("Element at (0,0):", twoDArrays[0][0])
	 fmt.Println("Element at (1,2):", twoDArrays[1][2])
	 fmt.Println("Element at (2,1):", twoDArrays[2][1])

	 //modifying elements
	 twoDArrays[0][0] = 10
	 twoDArrays[1][1] = 50
	 fmt.Println("Modified 2D Array:", twoDArrays)

	 //length of 2D array
	 fmt.Println("Number of rows in 2D array:", len(twoDArrays))
	 fmt.Println("Number of columns in 2D array:", len(twoDArrays[0]))
}