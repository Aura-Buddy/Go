package main

import "fmt"

//Defining a struct
type Test struct{
	Matrix [][]int
}

func main() {
	var a Test
	fmt.Println("The value of struct Test is: \n",a)
	a1:=Test{[][]int {{1},{2},{3}}}
	fmt.Println("The value of struct Test is: ", a1)

}
