package main

import "fmt"

func printArrays(matrix [][]int) {
	var i int
	for i = 0; i < 2; i++ {
		fmt.Println("The matrix is: \n", matrix[i])
	}
}

func main() {
	var a = [][]int {{1,3},{2}}
	printArrays(a)
}
