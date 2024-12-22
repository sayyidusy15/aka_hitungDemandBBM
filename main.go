package main

import (
	"fmt"
)

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

// ======  Implementasi Quick Sort Rekursif ======
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

//  ====== Implementasi Quick Sort Iteratif ======
type Stack struct {
	items [][2]int
}

func (s *Stack) Push(item [2]int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() [2]int {
	if len(s.items) == 0 {
		return [2]int{-1, -1}
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func QuickSortIterative(arr []FuelDemand, low, high int, isAscending bool) {
	stack := Stack{}
	stack.Push([2]int{low, high})

	for !stack.IsEmpty() {
		bounds := stack.Pop()
		low, high = bounds[0], bounds[1]

		if low < high {
			pivot := partitionIterative(arr, low, high, isAscending)

			if pivot-1 > low {
				stack.Push([2]int{low, pivot - 1})
			}
			if pivot+1 < high {
				stack.Push([2]int{pivot + 1, high})
			}
		}
	}
}

func partitionIterative(arr []FuelDemand, low, high int, isAscending bool) int {
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
	dataIterative := make([]FuelDemand, len(fuelDemands))
	copy(dataRecursive, fuelDemands)
	copy(dataIterative, fuelDemands)

	// Quick short section
	QuickSortRecursive(dataRecursive, 0, len(dataRecursive)-1, true)
	fmt.Println("\nData Quick Sort Reqursif:")
	printData(dataRecursive)

	// Quick short Iteratif
	QuickSortIterative(dataIterative, 0, len(dataIterative)-1, true)
	fmt.Println("\nHasil Quick Sort Iteratif")
	printData(dataIterative)

}
