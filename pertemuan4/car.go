package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) CalculateJarak() float64 {
	fuelPerMil := 1.5 // L/mill
	Jarak := c.FuelIn * fuelPerMil
	return Jarak
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 30.0,
	}

	Jarak := myCar.CalculateJarak()
	fmt.Printf("Mobil tipe %s dengan bahan bakar %.2f L dapat menempuh perkiraan jarak %.2f mill\n", myCar.Type, myCar.FuelIn, Jarak)
}
