package main

import "fmt"

func findMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Masukan 6 angka")
	for i := 0; i < 6; i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		fmt.Scan(&inputNumbers[i])
	}

	min, max := findMinMax(inputNumbers[:])

	// Referencing dengan pointer
	minPointer := &min
	maxPointer := &max

	// Dereferencing
	fmt.Println("Nilai minimum :", *minPointer)
	fmt.Println("Nilai maksimum :", *maxPointer)
}
