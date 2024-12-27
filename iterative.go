package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Struktur untuk menyimpan data daerah
type RegionDemand struct {
    Region string
    Demand int
}

// Quick Sort Iteratif dengan Optimasi
func quickSortIterative(data []RegionDemand) []RegionDemand {
    type rangeStack struct {
        start, end int
    }

    stack := []rangeStack{{0, len(data) - 1}}

    for len(stack) > 0 {
        currentRange := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        start, end := currentRange.start, currentRange.end
        if start >= end {
            continue
        }

        // Median-of-three pivot selection
        mid := start + (end-start)/2
        if data[start].Demand > data[mid].Demand {
            data[start], data[mid] = data[mid], data[start]
        }
        if data[mid].Demand > data[end].Demand {
            data[mid], data[end] = data[end], data[mid]
        }
        if data[start].Demand > data[mid].Demand {
            data[start], data[mid] = data[mid], data[start]
        }
        pivot := data[mid].Demand

        // Partition the array
        data[mid], data[end] = data[end], data[mid]
        low := start

        for i := start; i < end; i++ {
            if data[i].Demand < pivot {
                data[i], data[low] = data[low], data[i]
                low++
            }
        }

        data[low], data[end] = data[end], data[low]

        // Push smaller range first to minimize stack depth
        if low-1-start > end-low-1 {
            stack = append(stack, rangeStack{start, low - 1})
            stack = append(stack, rangeStack{low + 1, end})
        } else {
            stack = append(stack, rangeStack{low + 1, end})
            stack = append(stack, rangeStack{start, low - 1})
        }
    }

    return data
}

func printMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory Usage: %v KB\n", m.Alloc/1024)
}

func measurePerformance(data []RegionDemand, wg *sync.WaitGroup) {
    defer wg.Done()
    startTime := time.Now()

    printMemoryUsage() // Cetak penggunaan memori sebelum pengurutan

    quickSortIterative(data)

    duration := time.Since(startTime)

    printMemoryUsage() // Cetak penggunaan memori setelah pengurutan

    fmt.Printf("Data length: %d\n", len(data))
    fmt.Printf("Execution Time: %s\n\n", duration)
}

func main() {
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

    inputSizes := []int{1000, 10000, 50000, 100000}
    var wg sync.WaitGroup

    for _, targetLength := range inputSizes {
        fmt.Printf("\nTarget Length: %d\n", targetLength)

        initialData := data
        for len(initialData) < targetLength {
            remaining := targetLength - len(initialData)
            if remaining >= len(data) {
                initialData = append(initialData, data...)
            } else {
                initialData = append(initialData, data[:remaining]...)
            }
        }

        wg.Add(1)
        measurePerformance(initialData, &wg)
    }

    wg.Wait()
}
