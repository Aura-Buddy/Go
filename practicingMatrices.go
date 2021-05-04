package main

import "fmt"

func main() {
	var a = [3][4]int{{0,1,2,3},{4,5,6,7},{8,9,10,11}}

	fmt.Printf("The Matrix length is: %d\n", len(a))
	fmt.Printf("The Arrays length is: %d", len(a[1]))
}

