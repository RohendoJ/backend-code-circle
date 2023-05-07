package main

import "fmt"

func printFruits(name string, fruits []string) {
    fmt.Print("Hello, my name is ", name, ". My favorite fruits are: ")
    for _, fruit := range fruits {
        fmt.Print(fruit + " ")
    }
}

func main() {
    fruits := []string{"apple", "banana", "orange"}
    printFruits("John", fruits)
}
