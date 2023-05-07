package main

import "fmt"

func main() {
	for counter := 1; counter <= 20; counter++ {
		if counter%3 == 0 && counter%2 == 1 {
			fmt.Println(counter, "i love coding")
		} else if counter%3 == 0 {
			fmt.Println(counter, "code")
		} else if counter%2 == 0 {
			fmt.Println(counter, "circle")
		} else {
			fmt.Println(counter, "counter")
		}
	}
}
