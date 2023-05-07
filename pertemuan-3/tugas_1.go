package main

import "fmt"

func main() {	
	rectangleLength := 8
	rectangleWidth := 5

	triangleBase := 6.0
	triangleHeight := 7.0

	rectanglePerimeter := 2 * (rectangleLength + rectangleWidth)

	triangleArea := 0.5 * triangleBase * triangleHeight	

	fmt.Println("Perimeter of rectangle:", rectanglePerimeter)
	fmt.Println("Area of triangle:", triangleArea)
}