package main

import (
	"fmt"
)

func main() {
	// input := "1223456634127"
	// input := "76523752"
	input := "1234123"
	uniqueNumbers := findUniqueNumbers(input)

	fmt.Println("Angka yang muncul hanya 1 kali:", uniqueNumbers)
}

func findUniqueNumbers(input string) []int {
	counts := make(map[int]int)

	// Menghitung jumlah kemunculan masing-masing angka
	for _, char := range input {
		number := int(char - '0')
		counts[number]++
	}

	uniqueNumbers := []int{}

	// Menambahkan angka yang hanya muncul 1 kali ke dalam uniqueNumbers
	for number, count := range counts {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, number)
		}
	}

	return uniqueNumbers
}
