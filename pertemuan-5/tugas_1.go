package main

import "fmt"

func rectangleArea(length, width float64) (float64) {
	area := length * width
	return area
}

func rectanglePerimeter(length, width float64) (float64) {
	perimeter := 2 * (length + width)
	return perimeter
}

func cuboidVolume(length, width, height float64) (float64) {
	volume := length * width * height
	return volume
}

func main() {
	fmt.Println(rectangleArea(14, 2.5))
	fmt.Println(rectanglePerimeter(14, 2.5))
	fmt.Println(cuboidVolume(14, 2.5, 90.5))
}