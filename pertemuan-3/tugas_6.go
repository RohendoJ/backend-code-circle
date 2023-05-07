package main

import "fmt"

func main() {
	tahun := 2003

	if tahun >= 1944 && tahun <= 1964 {
		fmt.Print("boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Print("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Print("Generasi Y")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Print("Generasi Z")
	} else {
		fmt.Print("-")
	}
}