package main

import "fmt"

type FuelDemand struct {
	Region string
	Demand float64
}

// Data statis demand BBM
var fuelDemands = []FuelDemand{
	{"Jakarta", 10000},
	{"Surabaya", 8500},
	{"Bandung", 7200},
	{"Semarang", 6800},
	{"Medan", 6400},
	{"Makassar", 5900},
	{"Palembang", 5700},
}

// function to print the array
func printData(data []FuelDemand) {
	fmt.Println("\nDaftar Demand Bahan Bakar:")
	fmt.Println("==========================")
	for i, item := range data {
		fmt.Printf("%d. Region: %s, Demand: %.f\n", i+1, item.Region, item.Demand)
	}
	fmt.Println("==========================")
}

func main() {
	fmt.Println("Program Quick Sort Demand Bahan Bakar")
	fmt.Println("Data Awal:")
	printData(fuelDemands)
}
