package main

import "fmt"

// mengurutkan abjad a hingga z
// tiap abjad urutan kelipatan 5 diganti dengan tanda "%"
func wordplayOne() {
	for i := 'a'; i <= 'z'; i++ {
		if (i-'a'+1)%5 == 0 {
			fmt.Print("% ")
		} else {
			fmt.Printf("%c ", i)
		}
	}
	fmt.Println("")
}

func main() {
	fmt.Println("Selamat datang di WordPlay")
}
