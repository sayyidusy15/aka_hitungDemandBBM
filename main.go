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

// Implementasi Quick Sort Rekursif
func QuickSortRecursive(arr []FuelDemand, low, high int, isAscending bool) {
	if low < high {
		pi := partitionRecursive(arr, low, high, isAscending)
		QuickSortRecursive(arr, low, pi-1, isAscending)
		QuickSortRecursive(arr, pi+1, high, isAscending)
	}
}

func partitionRecursive(arr []FuelDemand, low, high int, isAscending bool) int {
	pivot := arr[high].Demand
	i := low - 1

	for j := low; j < high; j++ {
		if (isAscending && arr[j].Demand <= pivot) || (!isAscending && arr[j].Demand >= pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
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


	//menampilkan data pada array 
	dataRecursive := make([]FuelDemand, len(fuelDemands))
	copy(dataRecursive, fuelDemands)

	// Quick short section
	QuickSortRecursive(dataRecursive, 0, len(dataRecursive)-1, true)
	fmt.Println("\nData Setelah Quick Sort Ascending:")
	printData(dataRecursive)

}
