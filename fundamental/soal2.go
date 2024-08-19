package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var productName, expirationDateInput string
	var productPrice float64
	var stockQuantity int

	fmt.Print("Masukkan nama produk: ")
	fmt.Scanln(&productName)

	fmt.Print("Masukkan harga produk: ")
	fmt.Scanln(&productPrice)

	fmt.Print("Masukkan jumlah stok: ")
	fmt.Scanln(&stockQuantity)

	fmt.Print("Masukkan tanggal kadaluarsa (format YYYY-MM-DD): ")
	fmt.Scanln(&expirationDateInput)

	// ! Menjumlahkan harga Product !
	totalStockValue := productPrice * float64(stockQuantity)

	fmt.Println("\nInformasi Produk:")
	// ! Nama produk diformat Menjadi huruf kapital !
	fmt.Printf("Nama Produk: %s\n", strings.ToUpper(productName))

	// ! Harga produk diformat Menjadi mata uang Rupiah !
	fmt.Printf("Harga Produk: Rp %.2f\n", productPrice)

	// ! Jumlah stok dalam format dengan pemisah ribuan  !
	fmt.Printf("Jumlah Stok: %s\n", formatWithThousandSeparator(stockQuantity))

	// ! Tanggal kadaluarsa harus dalam format DD MMM YYYY !
	expirationDate, err := time.Parse("2006-01-02", expirationDateInput)
	if err != nil {
		// ! Jika terdapat nilai pada variable [err] maka akan mengembalikan nilai ( Format tanggal salah ) !
		fmt.Println("Format tanggal kadaluarsa salah.")
	} else {
		// ! Jika tidak ada nilai pada variable [err] maka akan mengembalikan Perubahan format Tanggal !
		fmt.Printf("Tanggal Kadaluarsa: %s\n", expirationDate.Format("02 Jan 2006"))
	}

	// Mengembalikan nilai total stok !
	fmt.Printf("Total Nilai Stok: Rp %.2f\n", totalStockValue)
}

// [!] formatWithThousandSeparator menambahkan pemisah ribuan pada jumlah stok
func formatWithThousandSeparator(number int) string {
	return fmt.Sprintf("%d", number)
}
