package main

import "fmt"

func main() {
	var kalimat = [...]string{"aku","dan","saya","sangat","senang","belajar","golang"}

	var kata []string
	for i := 2; i < len(kalimat); i++ {
		kata = append(kata, kalimat[i])
	}
	fmt.Println(kata)
}