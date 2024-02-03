package main

import "fmt"

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

func main() {
	fmt.Println("Selamat datang di WordPlay")
	wordplayOne()
}
