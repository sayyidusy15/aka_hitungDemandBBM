package main

import (
	"fmt"
	"time"
)

// Struktur untuk menyimpan data daerah
type RegionDemand struct {
	Region string
	Demand int
}

// Quick Sort Iteratif
func quickSortIterative(data []RegionDemand) []RegionDemand {
	stack := []struct {
		start, end int
	}{{0, len(data) - 1}}

	for len(stack) > 0 {
		n := len(stack) - 1
		start, end := stack[n].start, stack[n].end
		stack = stack[:n]

		if start >= end {
			continue
		}

		pivot := data[end].Demand
		low := start

		for i := start; i < end; i++ {
			if data[i].Demand < pivot {
				data[i], data[low] = data[low], data[i]
				low++
			}
		}

		data[low], data[end] = data[end], data[low]

		stack = append(stack, struct{ start, end int }{start, low - 1})
		stack = append(stack, struct{ start, end int }{low + 1, end})
	}

	return data
}

func main() {
	// Data simulasi
	data := []RegionDemand{
		{"Jakarta Pusat", 1500},
		{"Jakarta Selatan", 1400},
		{"Jakarta Utara", 1350},
		{"Bandung Barat", 1200},
		{"Bandung Timur", 1100},
		{"Surabaya Pusat", 1800},
		{"Surabaya Barat", 1700},
		{"Surabaya Timur", 1750},
		{"Yogyakarta Kota", 800},
		{"Yogyakarta Sleman", 850},
		{"Medan Kota", 1400},
		{"Medan Marelan", 1300},
		{"Semarang Tengah", 1300},
		{"Semarang Barat", 1250},
		{"Makassar Pusat", 1600},
		{"Makassar Timur", 1550},
		{"Denpasar Utara", 1100},
		{"Denpasar Selatan", 1050},
		{"Balikpapan Kota", 1700},
		{"Balikpapan Utara", 1650},
		{"Palembang Pusat", 900},
		{"Palembang Ilir Timur", 950},
		{"Pekanbaru Kota", 1000},
		{"Pekanbaru Tenayan", 1050},
		{"Banjarmasin Utara", 1400},
		{"Banjarmasin Selatan", 1350},
		{"Malang Kota", 1250},
		{"Malang Batu", 1200},
		{"Manado Kota", 1350},
		{"Manado Mapanget", 1400},
		{"Padang Kota", 1450},
		{"Padang Selatan", 1350},
		{"Pontianak Kota", 1150},
		{"Pontianak Utara", 1200},
		{"Batam Kota", 950},
		{"Batam Nongsa", 900},
		{"Bogor Kota", 1550},
		{"Bogor Barat", 1500},
		{"Cirebon Kota", 850},
		{"Cirebon Tengah", 800},
		{"Solo Kota", 700},
		{"Solo Sukoharjo", 750},
		{"Balikpapan Baru", 1600},
		{"Depok Margonda", 1400},
		{"Depok Cimanggis", 1450},
		{"Tangerang Kota", 1350},
		{"Tangerang Selatan", 1250},
		{"Bekasi Timur", 1100},
		{"Bekasi Barat", 1150},
		{"Karawang Kota", 1200},
		{"Karawang Timur", 1250},
		{"Cilegon Kota", 950},
		{"Cilegon Merak", 1000},
	}

	numIterations := 10 // Number of times to duplicate the array
	for i := 0; i < numIterations; i++ {
		data = append(data, data...) // Append all current elements to the array
	}

	// Mengukur waktu eksekusi
	startTime := time.Now()

	// Pengurutan
	// sortedData := quickSortIterative(data)
	quickSortIterative(data)

	duration := time.Since(startTime)
	// Output hasil
	// for _, region := range sortedData {
	// 	fmt.Printf("Region: %s, Demand: %d\n", region.Region, region.Demand)
	// }

	fmt.Printf("Data length: %d\n", len(data))
	fmt.Printf("Execution Time: %s\n", duration)
}
