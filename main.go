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

// menyebutkan factory sebanyak 9 kali
// tiap factory urutan kelipatan 2 diubah menggunakan uppercase
func wordplayTwo() {
	for i := 1; i <= 9; i++ {
		word := "factory"
		if i%2 == 0 {
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
