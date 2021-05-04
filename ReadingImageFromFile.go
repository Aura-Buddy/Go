package main

import(
	"fmt"
	//"image"
	"image/png"
	"os"
)

func main() {
	existingImageFile,err:=os.Open("test1.png")
	if err != nil {
		fmt.Println("Error opening test1.png")
	}
	defer  existingImageFile.Close()
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		fmt.Println("Ohh lordy lordy")
	}

	fmt.Println(loadedImage)
}
