package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Masukkan Angka :")
	var numb int
	fmt.Scan(&numb)
	fmt.Println(bilanganKelipatan(numb, 7))
}
func bilanganKelipatan(n int, k int) (hasil string) {
	if n < k {
		return strconv.Itoa(n) + " bukan bilangan kelipatan " + strconv.Itoa(k)
		if n == k || n%k == 0 {
			return strconv.Itoa(n) + " adalah bilangan kelipatan " + strconv.Itoa(k)
		}
	} else {
		return strconv.Itoa(n) + " adalah bilangan kelipatan " + strconv.Itoa(k)
	}
	return
}
