package main

import (
	"image"
	"fmt"
)

func main() {
	myImage:= image.NewRGBA(image.Rect(0,0,10,4))
	myImage.Pix[0] = 255
	myImage.Pix[1] = 0
	myImage.Pix[2] = 0
	myImage.Pix[3] = 255

	fmt.Println(myImage.Pix)
	fmt.Println(myImage.Stride)
}
