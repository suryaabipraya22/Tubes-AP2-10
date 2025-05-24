package main

import (
	"fmt"
)

const NMAX = 100

type Pengeluaran struct {
	kategori string
	jumlah   int
}

type arrPengeluaran [NMAX]Pengeluaran

func main() {
	var data arrPengeluaran
	var count int
	var budget int
	var pilihan int

	fmt.Print("Masukkan total budget perjalanan Anda: ")
	fmt.Scanln(&budget)

	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahData(&data, &count)
		case 2:
			ubahData(&data, count)
		case 3:
			hapusData(&data, &count)
		case 4:
			tampilkanData(data, count)
		case 5:
			hitungTotalDanSaran(data, count, budget)
		case 6:
			cariDataKategori(data, count)
		case 7:
			urutkanData(&data, count)
		case 8:
			tampilkanLaporan(data, count, budget)
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi kami!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║              MENU PENGELOLAAN BUDGET           ║")
	fmt.Println("╠════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Tambah Pengeluaran                          ║")
	fmt.Println("║ 2. Ubah Pengeluaran                            ║")
	fmt.Println("║ 3. Hapus Pengeluaran                           ║")
	fmt.Println("║ 4. Tampilkan Seluruh Pengeluaran               ║")
	fmt.Println("║ 5. Hitung Total & Saran Penghematan            ║")
	fmt.Println("║ 6. Cari Pengeluaran (Sequential Search)        ║")
	fmt.Println("║ 7. Urutkan Pengeluaran                         ║")
	fmt.Println("║ 8. Tampilkan Laporan                           ║")
	fmt.Println("║ 9. Keluar                                      ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Print("▶ Pilih menu: ")
}

func tambahData(data *arrPengeluaran, count *int) {
	if *count < NMAX {
		var pilihan int
		var jml int
		var kat string

		for pilihan < 1 || pilihan > 5 {
			fmt.Println("Pilih kategori pengeluaran:")
			fmt.Println("1. Transportasi")
			fmt.Println("2. Hotel")
			fmt.Println("3. Makanan")
			fmt.Println("4. Hiburan")
			fmt.Println("5. Oleh-oleh")
			fmt.Print("Pilihan (1-5): ")
			fmt.Scanln(&pilihan)

			if pilihan < 1 || pilihan > 5 {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		}

		switch pilihan {
		case 1:
			kat = "Transportasi"
		case 2:
			kat = "Akomodasi"
		case 3:
			kat = "Makanan"
		case 4:
			kat = "Hiburan"
		case 5:
			kat = "Oleh-oleh"
		}

		fmt.Print("Masukkan jumlah pengeluaran: ")
		fmt.Scanln(&jml)

		data[*count] = Pengeluaran{kat, jml}
		*count++
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Data penuh. Tidak bisa menambah lagi.")
	}
}

func ubahData(data *arrPengeluaran, count int) {
	var index int
	fmt.Print("Masukkan indeks pengeluaran yang ingin diubah: ")
	fmt.Scanln(&index)

	if index >= 0 && index < count {
		var pilihan int
		for pilihan < 1 || pilihan > 5 {
			fmt.Println("Pilih kategori baru:")
			fmt.Println("1. Transportasi")
			fmt.Println("2. Hotel")
			fmt.Println("3. Makanan")
			fmt.Println("4. Hiburan")
			fmt.Println("5. Oleh-oleh")
			fmt.Print("Pilihan (1-5): ")
			fmt.Scanln(&pilihan)

			if pilihan < 1 || pilihan > 5 {
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		}

		switch pilihan {
		case 1:
			data[index].kategori = "Transportasi"
		case 2:
			data[index].kategori = "Akomodasi"
		case 3:
			data[index].kategori = "Makanan"
		case 4:
			data[index].kategori = "Hiburan"
		case 5:
			data[index].kategori = "Oleh-oleh"
		}

		fmt.Print("Masukkan jumlah baru: ")
		fmt.Scanln(&data[index].jumlah)
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusData(data *arrPengeluaran, count *int) {
	var index int
	fmt.Print("Masukkan indeks pengeluaran yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index >= 0 && index < *count {
		for i := index; i < *count-1; i++ {
			data[i] = data[i+1]
		}
		*count--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func tampilkanData(data arrPengeluaran, count int) {
	if count == 0 {
		fmt.Println("Belum ada data pengeluaran.")
	} else {
		fmt.Println("Daftar Pengeluaran:")
		for i := 0; i < count; i++ {
			fmt.Printf("%d. %s - %d\n", i, data[i].kategori, data[i].jumlah)
		}
	}
}

func hitungTotalDanSaran(data arrPengeluaran, count int, budget int) {
	total := 0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}

	fmt.Printf("Total pengeluaran: %d\n", total)
	if total > budget {
		fmt.Printf("Anda melebihi budget sebesar %d. Kurangi pengeluaran Anda!\n", total-budget)
	} else {
		fmt.Printf("Masih ada sisa budget: %d. Anda cukup hemat!\n", budget-total)
	}
}

func cariDataKategori(data arrPengeluaran, count int) {
	var kat string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scanln(&kat)
	found := false

	for i := 0; i < count; i++ {
		if data[i].kategori == kat {
			fmt.Printf("Ditemukan di indeks %d: %s - %d\n", i, data[i].kategori, data[i].jumlah)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

func urutkanData(data *arrPengeluaran, count int) {
	var pilihan int
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort berdasarkan jumlah")
	fmt.Println("2. Insertion Sort berdasarkan kategori")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
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

func tampilkanLaporan(data arrPengeluaran, count int, budget int) {
	if count == 0 {
		fmt.Println("Tidak ada data pengeluaran.")
		return
	}

	fmt.Println("=== LAPORAN PENGELUARAN ===")
	tampilkanData(data, count)

	total := 0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}
	selisih := total - budget
	if selisih < 0 {
		selisih = -selisih
	}

	fmt.Printf("Total Pengeluaran: %d\n", total)
	fmt.Printf("Budget: %d\n", budget)
	fmt.Printf("Selisih Budget dan Pengeluaran: %d\n", selisih)
}
