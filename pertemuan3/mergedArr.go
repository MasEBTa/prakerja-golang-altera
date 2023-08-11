package main

import "fmt"

func main() {
	array1 := []string{"John", "Jane", "Alice"}
	array2 := []string{"Bob", "Jane", "Charlie"}

	merged := mergeArrays(array1, array2)
	fmt.Println("Array setelah digabung dan dihapus elemen yang sama:", merged)
}

func mergeArrays(arr1, arr2 []string) []string {
	merged := append(arr1, arr2...) // Menggabungkan array menggunakan variadic argument
	result := make([]string, 0)     // Membuat slice kosong untuk hasil akhir

	// Loop melalui array yang digabungkan
	for _, item := range merged {
		if !contains(result, item) {
			result = append(result, item) // Menambahkan item ke hasil jika belum ada
		}
	}

	return result
}

func contains(arr []string, item string) bool {
	for _, value := range arr {
		if value == item {
			return true
		}
	}
	return false
}
