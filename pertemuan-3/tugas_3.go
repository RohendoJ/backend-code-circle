package main

import "fmt"

func main() {
	tanggal := 4
	bulan := 6
	tahun := 2003

	fmt.Print(tanggal)

	switch bulan {
	case 1:
		fmt.Print("Januari")
	case 2:
		fmt.Print("Februari")
	case 3:
		fmt.Print("Maret")
	case 4:
		fmt.Print("April")
	case 5:
		fmt.Print("Mei")
	case 6:
		fmt.Print("Juni")
	case 7:
		fmt.Print("Juli")
	case 8:
		fmt.Print("Agustus")
	case 9:
		fmt.Print("September")
	case 10:
		fmt.Print("Oktober")
	case 11:
		fmt.Print("November")
	case 12:
		fmt.Print("Desember")
	}

	fmt.Print(tahun)
}