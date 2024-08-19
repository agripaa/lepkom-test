package main

import (
	"fmt"
	"strings"
)

func main() {
	// ! Meminta input kata dari pengguna !
	var inputWord string
	fmt.Print("Masukkan sebuah kata: ")
	fmt.Scanln(&inputWord)

	// ! Memproses kata untuk memisahkan huruf vokal dan konsonan !
	vowels, consonants := separateVowelsAndConsonants(inputWord)

	// ! Menampilkan hasil pemisahan huruf vokal !
	fmt.Printf("Huruf Vokal: %s\n", vowels)

	// ! Menampilkan hasil pemisahan huruf konsonan !
	fmt.Printf("Huruf Konsonan: %s\n", consonants)
}

// [!] separateVowelsAndConsonants memisahkan huruf vokal dan konsonan dari kata yang diberikan
func separateVowelsAndConsonants(word string) (string, string) {
	vowels := "aeiouAEIOU"
	var vowelLetters, consonantLetters string

	// ! Iterasi melalui setiap huruf dalam kata !
	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			// ! Jika karakter merupakan huruf vokal, tambahkan ke vowelLetters !
			vowelLetters += string(char)
		} else if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			// ! Jika karakter merupakan huruf konsonan, tambahkan ke consonantLetters !
			consonantLetters += string(char)
		}
	}

	// Mengembalikan hasil pemisahan huruf vokal dan konsonan
	return vowelLetters, consonantLetters
}
