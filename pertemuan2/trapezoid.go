package main

import (
	"fmt"
)

func main() {
	// variable
	var a, b, t int
	// scan
	fmt.Println("Masukkan Panjang sisi Atas :")
	fmt.Scan(&a)
	fmt.Println("Masukkan Panjang sisi Bawah :")
	fmt.Scan(&b)
	fmt.Println("Masukkan Panjang sisi Tinggi :")
	fmt.Scan(&t)
	// luas
	fmt.Println("luas Trapesium Anda :", luasTrapezoid(a, b, t))
}

// function luas
func luasTrapezoid(atas, bawah, tinggi int) (hasil float64) {
	aDaNb := float64(atas + bawah)
	hasil = 0.5 * aDaNb * float64(tinggi)
	return
}
