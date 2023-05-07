package main

import "fmt"

func introduce(name, gender, occupation, age string) string {
	biodata := "Name: " + name + "\n" + "Gender: " + gender + "\n" + "Occupation: " + occupation + "\n" + "Age: " + age
	return biodata
}

func main() {
	fmt.Println(introduce("Rohendo", "Male", "Student", "19"))
}