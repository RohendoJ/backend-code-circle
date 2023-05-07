package main

import "fmt"

func main() {
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam","Buncis","Kangkung","Kubis","Seledri","Tauge","Timun")

	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d. %s\n", i+1, sayuran[i])
	}
}