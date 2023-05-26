package main

import "fmt"

func main() {
	var length, width float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scanln(&length)

	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	area := length * width
	perimeter := 2 * (length + width)

	fmt.Printf("Area of the rectangle: %.2f\n", area)
	fmt.Printf("Perimeter of the rectangle: %.2f\n", perimeter)
}
