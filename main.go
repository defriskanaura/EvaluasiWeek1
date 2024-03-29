package main

import (
	"fmt"
	"strings"
)

// mengurutkan abjad f hingga y
// tiap abjad urutan kelipatan 2 diganti dengan tanda "re"
func wordplayOne() {
	for i := 'f'; i <= 'y'; i++ {
		if (i-'f'+1)%2 == 0 {
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

// menyebutkan sabang sebanyak 20 kali
// tiap sabang urutan kelipatan 5 diubah menggunakan uppercase
func wordplayThree() {
	for i := 1; i <= 9; i++ {
		word := "SABANG-MERAUKE"
		if i%2 == 0 {
			fmt.Printf("%s ", strings.ToLower(word))
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
	wordplayThree()
}
