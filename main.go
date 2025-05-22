package main

import (
	"fmt"
	"math"
)

const NMAX = 100 // Maksimum jumlah data pengeluaran yang bisa disimpan

// Struktur data pengeluaran
type Pengeluaran struct {
	kategori string
	jumlah   float64
}

var data [NMAX]Pengeluaran // Array untuk menyimpan data pengeluaran
var count int              // Jumlah data yang sedang tersimpan
var budget float64         // Budget total yang direncanakan

func main() {
	var pilihan int

	// Input budget awal dari pengguna
	fmt.Print("Masukkan total budget perjalanan Anda: ")
	fmt.Scanln(&budget)

	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			hapusData()
		case 4:
			tampilkanData()
		case 5:
			hitungTotalDanSaran()
		case 6:
			cariDataKategori()
		case 7:
			urutkanData()
		case 8:
			tampilkanLaporan()
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menampilkan menu utama
func tampilkanMenu() {
	fmt.Println("\n=== MENU PENGELOLAAN BUDGET TRAVELING ===")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Ubah Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Tampilkan Seluruh Pengeluaran")
	fmt.Println("5. Hitung Total & Saran Penghematan")
	fmt.Println("6. Cari Pengeluaran (Sequential Search)")
	fmt.Println("7. Urutkan Pengeluaran")
	fmt.Println("8. Tampilkan Laporan")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu: ")
}

// Menambahkan pengeluaran baru
func tambahData() {
	if count < NMAX {
		var kat string
		var jml float64
		fmt.Print("Masukkan kategori pengeluaran: ")
		fmt.Scanln(&kat)
		fmt.Print("Masukkan jumlah pengeluaran: ")
		fmt.Scanln(&jml)

		data[count] = Pengeluaran{kat, jml}
		count++
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Data penuh. Tidak bisa menambah lagi.")
	}
}

// Mengubah data pengeluaran berdasarkan indeks
func ubahData() {
	var index int
	fmt.Print("Masukkan indeks pengeluaran yang ingin diubah: ")
	fmt.Scanln(&index)

	if index >= 0 && index < count {
		fmt.Print("Masukkan kategori baru: ")
		fmt.Scanln(&data[index].kategori)
		fmt.Print("Masukkan jumlah baru: ")
		fmt.Scanln(&data[index].jumlah)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// Menghapus data pengeluaran dengan menggeser array
func hapusData() {
	var index int
	fmt.Print("Masukkan indeks pengeluaran yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index >= 0 && index < count {
		for i := index; i < count-1; i++ {
			data[i] = data[i+1]
		}
		count--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

// Menampilkan seluruh pengeluaran yang telah ditambahkan
func tampilkanData() {
	if count == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar Pengeluaran:")
		for i := 0; i < count; i++ {
			fmt.Printf("%d. %s - %.2f\n", i, data[i].kategori, data[i].jumlah)
		}
	}
}

// Menghitung total pengeluaran dan menampilkan saran penghematan
func hitungTotalDanSaran() {
	total := 0.0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}

	fmt.Printf("Total pengeluaran: %.2f\n", total)
	if total > budget {
		fmt.Printf("Anda melebihi budget sebesar %.2f. Kurangi pengeluaran Anda!\n", total-budget)
	} else {
		fmt.Printf("Masih ada sisa budget: %.2f. Anda cukup hemat!\n", budget-total)
	}
}

// Mencari pengeluaran berdasarkan kategori dengan Sequential Search
func cariDataKategori() {
	var kat string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scanln(&kat)
	found := false

	for i := 0; i < count; i++ {
		if data[i].kategori == kat {
			fmt.Printf("Ditemukan di indeks %d: %s - %.2f\n", i, data[i].kategori, data[i].jumlah)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Mengurutkan data pengeluaran dengan Selection dan Insertion Sort
func urutkanData() {
	var pilihan int
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort berdasarkan jumlah")
	fmt.Println("2. Insertion Sort berdasarkan kategori")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		// Selection Sort by jumlah
		for i := 0; i < count-1; i++ {
			min := i
			for j := i + 1; j < count; j++ {
				if data[j].jumlah < data[min].jumlah {
					min = j
				}
			}
			data[i], data[min] = data[min], data[i]
		}
		fmt.Println("Data diurutkan berdasarkan jumlah (Selection Sort).")
	} else if pilihan == 2 {
		// Insertion Sort by kategori
		for i := 1; i < count; i++ {
			temp := data[i]
			j := i - 1
			for j >= 0 && data[j].kategori > temp.kategori {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
		fmt.Println("Data diurutkan berdasarkan kategori (Insertion Sort).")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// Menampilkan laporan: daftar, total, dan selisih dengan budget
func tampilkanLaporan() {
	if count == 0 {
		fmt.Println("Tidak ada data pengeluaran.")
		return
	}

	fmt.Println("=== LAPORAN PENGELUARAN ===")
	tampilkanData()

	total := 0.0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}
	fmt.Printf("Total Pengeluaran: %.2f\n", total)
	fmt.Printf("Budget: %.2f\n", budget)
	fmt.Printf("Selisih Budget dan Pengeluaran: %.2f\n ", math.Abs(budget-total))
}
