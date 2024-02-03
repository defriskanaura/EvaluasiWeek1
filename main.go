package main

import (
	"fmt"
	"strings"
)

// mengurutkan abjad e hingga x
// tiap abjad urutan kelipatan 2 diganti dengan tanda "re"
func wordplayOne() {
	for i := 'e'; i <= 'x'; i++ {
		if (i-'e'+1)%2 == 0 {
			fmt.Print("re ")
		} else {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println("")
}

// menyebutkan bee sebanyak 15 kali
// tiap bee urutan kelipatan 3 diubah menggunakan uppercase
func wordplayTwo() {
	for i := 1; i <= 15; i++ {
		word := "bee"
		if i%3 == 0 {
			// Ubah huruf pertama menjadi uppercase
			fmt.Printf("%s ", strings.ToUpper(word))
		} else {
			fmt.Printf("%s ", word)
		}
	}
	fmt.Println("")
}

func main() {
	fmt.Println("Selamat datang di WordPlay")
	wordplayOne()
	wordplayTwo()
}
