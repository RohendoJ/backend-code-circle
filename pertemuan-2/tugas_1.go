package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operation string
	var result float64

	fmt.Println("enter the first number : ")
	fmt.Scan(&num1)
	fmt.Println("enter the second number : ")
	fmt.Scan(&num2)
	fmt.Println("input operation [+,-.*,/] : ")
	fmt.Scan(&operation)

	if operation == "+" {
		result = num1 + num2 
		fmt.Print("output : ",result)
	} else if operation == "-" {
		result = num1 - num2 
		fmt.Print("output : ",result)
	} else if operation == "*" {
		result = num1 * num2 
		fmt.Print("output : ",result)
	} else if operation == "/" && num1 == 0 || num2 == 0 {
		fmt.Print("Error: cannot divide by zero")
	} else if operation == "/" {
		result = num1 / num2 
		fmt.Print("output : ",result)
	} else {
		fmt.Println("operation error")
	}

}