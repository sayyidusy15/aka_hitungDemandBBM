# Quick Sort untuk Analisis Permintaan Wilayah

Proyek ini menunjukkan implementasi algoritma Quick Sort menggunakan **Go** untuk mengurutkan data simulasi permintaan bahan bakar (BBM) di berbagai wilayah di Indonesia. Pengurutan ini bertujuan untuk mengidentifikasi wilayah dengan permintaan tertinggi hingga terendah guna mengoptimalkan perencanaan distribusi.

Nama Anggota:
- Sayyidusy Syauqi AL Ghiffari - 103012480033
- Muhammad Fathan Akbar - 103012480042

## Fitur
- **Quick Sort Rekursif**: Implementasi sederhana dan mudah dipahami.
- **Quick Sort Iteratif**: Manajemen memori yang efisien menggunakan stack manual.
- **Data Terstruktur**: Mengurutkan data yang direpresentasikan sebagai pasangan wilayah-permintaan.

## Daftar Isi
1. [Latar Belakang](#latar-belakang)
2. [Persyaratan](#persyaratan)
3. [Instalasi](#instalasi)
4. [Penggunaan](#penggunaan)
5. [Contoh](#contoh)
6. [Perbandingan](#perbandingan)
7. [Lisensi](#lisensi)

## Latar Belakang
Permintaan bahan bakar berbeda-beda di setiap wilayah di Indonesia akibat populasi, aktivitas industri, dan infrastruktur. Untuk memastikan distribusi yang efisien, wilayah harus diurutkan berdasarkan permintaan bahan bakarnya. Proyek ini mensimulasikan data tersebut dan mengimplementasikan algoritma Quick Sort untuk menyelesaikan masalah ini secara efisien.

## Persyaratan
- Go 1.16 atau yang lebih baru

## Instalasi
1. Clone repository:
   ```bash
   git clone https://github.com/your-username/quick-sort-bbm.git
   cd quick-sort-bbm
   ```
2. Bangun aplikasi:
   ```bash
   go build iterative.go
   go build recursive.go
   ```

## Penggunaan
Jalankan aplikasi untuk melihat data permintaan yang telah diurutkan:
```bash
./iterative
./recursive
```

### Input Data
Data didefinisikan dalam kode sebagai array struktur `RegionDemand`. Modifikasi variabel `data` di fungsi `main` untuk menguji dengan dataset yang berbeda.

Contoh:
```go
[]RegionDemand{
    {"Jakarta", 1500},
    {"Bandung", 1200},
    {"Surabaya", 1800},
    {"Yogyakarta", 800},
    {"Medan", 1400},
}
```

## Contoh
### Output Quick Sort Rekursif
Untuk data input:
```go
[]RegionDemand{
    {"Jakarta", 1500},
    {"Bandung", 1200},
    {"Surabaya", 1800},
    {"Yogyakarta", 800},
    {"Medan", 1400},
}
```

Output yang telah diurutkan:
```
Region: Yogyakarta, Demand: 800
Region: Bandung, Demand: 1200
Region: Medan, Demand: 1400
Region: Jakarta, Demand: 1500
Region: Surabaya, Demand: 1800
```

### Output Quick Sort Iteratif
Hasil untuk input yang sama tetap konsisten dengan pendekatan rekursif.

## Perbandingan
| **Kriteria**           | **Rekursif**               | **Iteratif**                |
|-------------------------|----------------------------|-----------------------------|
| **Kemudahan Implementasi** | Sederhana dan langsung     | Lebih kompleks              |
| **Penggunaan Memori**   | O(log n) untuk pemanggilan rekursif | O(1) dengan stack manual   |
| **Efisiensi**          | Terbaik untuk dataset kecil | Lebih baik untuk dataset besar |

## Lisensi
Proyek ini dilisensikan di bawah Lisensi MIT. Lihat file LICENSE untuk detail.
