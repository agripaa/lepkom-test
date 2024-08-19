package main

import (
	"fmt"
	"time"
)

func main() {
	// ! Inisialisasi Variable !
	var integerInput int
	var floatInput float64
	var dateInput string

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scanln(&integerInput)

	fmt.Print("Masukkan bilangan pecahan: ")
	fmt.Scanln(&floatInput)

	fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&dateInput)

	fmt.Println("\nBilangan bulat:")
	// ~ Konversi dan tampilkan bilangan bulat ke dalam berbagai format ~
	fmt.Printf("Desimal: %d\n", integerInput)
	fmt.Printf("Oktal: %o\n", integerInput)
	fmt.Printf("Heksadesimal: %X\n", integerInput)

	fmt.Println("\nBilangan pecahan:")
	fmt.Printf("Dua angka di belakang koma: %.2f\n", floatInput)

	fmt.Println("\nTanggal:")
	// ! Memberi validasi apakah Tanggal yang di inputkan sesuai atau tidak !
	date, err := time.Parse("2006-01-02", dateInput)

	if err != nil {
		// ! Jika terdapat nilai pada variable [err] maka akan mengembalikan nilai ( Format tanggal salah ) !
		fmt.Println("Format tanggal salah.")
	} else {
		// ! Jika tidak ada nilai pada variable [err] maka akan mengembalikan Perubahan format Tanggal tersebut !
		fmt.Println("Format DD-MM-YYYY:", date.Format("02-01-2006"))
	}
}
