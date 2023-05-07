package main

import "fmt"

func main() {
	var satuan = map[string]int{
		"panjang": 7,
		"lebar": 4,
		"tinggi": 6,
	}

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Printf("volume balok = %d", volume)
}