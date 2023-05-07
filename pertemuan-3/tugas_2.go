package main

import "fmt"

func main() {
	john := 80

	if john >= 80 {	
		fmt.Print("john A ")
	} else if john >= 70 && john < 80 {
		fmt.Print("john B ")
	} else if john >= 60 && john < 70 {
		fmt.Print("john C ")
	} else if john >= 50 && john < 60 {
		fmt.Print("john D ")
	} else if john < 50 {
		fmt.Print("john E ")
	} else {
		fmt.Print("ulang")
	}
	
	doe := 50

	if doe >= 80 {	 
		fmt.Print("doe A ")
	} else if doe >= 70 && doe < 80 {
		fmt.Print("doe B ")
	} else if doe >= 60 && doe < 70 {
		fmt.Print("doe C ")
	} else if doe >= 50 && doe < 60 {
		fmt.Print("doe D ")
	} else if doe < 50 {
		fmt.Print("doe E ")
	} else {
		fmt.Print("ulang")
	}
}