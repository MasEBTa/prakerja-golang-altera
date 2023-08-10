package main

import "fmt"

func main() {
	fmt.Println("Masukkan Angka :")
	var numb int
	fmt.Scan(&numb)
	bilanganPrima(numb)
}
func bilanganPrima(n int) {
	if n > 1 {
		for i := 2; i <= n; i++ {
			if n/i == 1 {
				fmt.Printf("%d adalah bilangan prima", n)
				return
			} else if n%i == 0 {
				fmt.Printf("%d bukan bilangan prima", n)
				return
			}
		}
	} else {
		fmt.Printf("%d bukan bilangan prima", n)
		return
	}
}
