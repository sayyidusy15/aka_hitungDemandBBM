package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FuelDemand struct {
	Region string
	Demand float64
}

// Fungsi untuk generate data random
func generateRandomData(size int) []FuelDemand {
	regions := []string{"Jakarta", "Surabaya", "Bandung", "Semarang", "Medan", "Makassar", "Palembang", 
		"Yogyakarta", "Malang", "Denpasar", "Balikpapan", "Manado", "Padang", "Pekanbaru", "Banjarmasin"}
	
	data := make([]FuelDemand, size)
	for i := 0; i < size; i++ {
		randomRegion := regions[rand.Intn(len(regions))]
		randomDemand := rand.Float64()*10000 + 1000 // Random demand antara 1000-11000
		data[i] = FuelDemand{randomRegion, randomDemand}
	}
	return data
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

// function to print the array (optional untuk debugging)
func printData(data []FuelDemand) {
	fmt.Println("\nDaftar Demand Bahan Bakar:")
	fmt.Println("==========================")
	for i, item := range data {
		fmt.Printf("%d. Region: %s, Demand: %.2f\n", i+1, item.Region, item.Demand)
	}
	fmt.Println("==========================")
}

func main() {
	// Set seed untuk random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random data
	dataSize := 20 // Ubah sesuai kebutuhan
	randomData := generateRandomData(dataSize)

	// Copy data untuk kedua metode sorting
	dataRecursive := make([]FuelDemand, len(randomData))
	dataIterative := make([]FuelDemand, len(randomData))
	copy(dataRecursive, randomData)
	copy(dataIterative, randomData)

	fmt.Printf("\nMelakukan pengurutan untuk %d data...\n", dataSize)

	// Quick sort Rekursif dengan pengukuran waktu
	startRecursive := time.Now()
	QuickSortRecursive(dataRecursive, 0, len(dataRecursive)-1, true)
	durationRecursive := time.Since(startRecursive)
	
	// Quick sort Iteratif dengan pengukuran waktu
	startIterative := time.Now()
	QuickSortIterative(dataIterative, 0, len(dataIterative)-1, true)
	durationIterative := time.Since(startIterative)

	// Menampilkan hasil pengukuran waktu dalam microsecond
	fmt.Printf("\nHasil Pengukuran Waktu untuk %d data:\n", dataSize)
	fmt.Printf("Quick Sort Rekursif: %v microsecond\n", durationRecursive.Microseconds())
	fmt.Printf("Quick Sort Iteratif: %v microsecond\n", durationIterative.Microseconds())

	fmt.Println("\nKeterangan satuan waktu:")
	fmt.Println("1 detik = 1000 millisecond = 1.000.000 microsecond")
	fmt.Println("- microsecond (µs) = sepersejuta detik")
	fmt.Println("- millisecond (ms) = seperseribu detik")
	fmt.Println("- second (s) = satu detik")

	// Uncomment baris di bawah ini jika ingin melihat data hasil sorting
	fmt.Println("\nHasil Quick Sort Rekursif:")
	printData(dataRecursive)
	fmt.Println("\nHasil Quick Sort Iteratif:")
	printData(dataIterative)
}