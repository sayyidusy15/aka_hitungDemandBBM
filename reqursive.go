package main

import (
    "fmt"
    "runtime"
    "time"
)

// Struktur untuk menyimpan data daerah
type RegionDemand struct {
    Region string
    Demand int
}

// Quick Sort Rekursif
func quickSortRecursive(data []RegionDemand) []RegionDemand {
    if len(data) <= 1 {
        return data
    }

    pivot := data[len(data)/2].Demand
    left := []RegionDemand{}
    right := []RegionDemand{}
    middle := []RegionDemand{}

    for _, val := range data {
        if val.Demand < pivot {
            left = append(left, val)
        } else if val.Demand > pivot {
            right = append(right, val)
        } else {
            middle = append(middle, val)
        }
    }

    left = quickSortRecursive(left)
    right = quickSortRecursive(right)

    return append(append(left, middle...), right...)
}

func printMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory Usage: %v KB\n", m.Alloc/1024)
}

func measurePerformance(data []RegionDemand) {
    // Mengukur waktu eksekusi
    startTime := time.Now()

    printMemoryUsage() // Cetak penggunaan memori sebelum pengurutan

    // Pengurutan
    quickSortRecursive(data)

    duration := time.Since(startTime)

    printMemoryUsage() // Cetak penggunaan memori setelah pengurutan

    // Output hasil
    fmt.Printf("Data length: %d\n", len(data))
    fmt.Printf("Execution Time: %s\n", duration)
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

    inputSizes := []int{1000, 10000, 50000, 100000, 1000000} // Variasi ukuran input

    for _, targetLength := range inputSizes {
        fmt.Printf("\nTarget Length: %d\n", targetLength)

        // Membuat data baru sesuai target length
        initialData := data
        for len(initialData) < targetLength {
            remaining := targetLength - len(initialData)
            if remaining >= len(data) {
                initialData = append(initialData, data...)
            } else {
                initialData = append(initialData, data[:remaining]...)
            }
        }

        measurePerformance(initialData)
    }
}
