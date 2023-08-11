package main

import (
	"fmt"
)

func main() {
	slice := []string{"apple", "banana", "apple", "orange", "apple", "grape"}

	stringToCount := "apple"
	// stringToCount := "papaya"
	// stringToCount := "banana"
	count := countStringInSlice(slice, stringToCount)

	fmt.Printf("Jumlah kemunculan '%s' dalam slice: %d\n", stringToCount, count)
}

func countStringInSlice(slice []string, target string) int {
	count := 0
	for _, s := range slice {
		if s == target {
			count++
		}
	}
	return count
}
