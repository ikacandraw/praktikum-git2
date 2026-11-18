package main

import "fmt"

func main() {
	fmt.Println("------------PROGRAM PARKIR-----------------")
	fmt.Println("Parkir Undiknas")
	jam := 0
	var totaljam, total int

	for jam != -1 {
		fmt.Print("Masukkan jam: ")
		fmt.Scan(&jam)
		totaljam += jam

	}
	total = 5000 + (totaljam-1)*2000
	fmt.Println("Jumlah jam parkir: ", totaljam)
	fmt.Println("Total bayar: ", total)

}
