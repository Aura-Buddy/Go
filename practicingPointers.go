package main

import "fmt"

func passingPointers(i *int) {
	fmt.Printf("The value of i is: %d\nIt's held in memory location: %d", *i, i)
}

func main() {
	i := [3]int{42,54,87}
	p:=&i[1]
	passingPointers(p)
}
