package main

import "fmt"

const NMAX int = 100

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

	// Isi data dummy (hapus/comment baris ini jika tidak ingin menggunakan dummy)
	isiDataDummy(&data, &count)

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
			fmt.Println("╔═══════════════════════════════════════╗")
			fmt.Println("║     TERIMA KASIH TELAH MENGGUNAKAN    ║")
			fmt.Println("║       APLIKASI PENGELOLA BUDGET!      ║")
			fmt.Println("╠═══════════════════════════════════════╣")
			fmt.Println("║        Semoga perjalanan Anda         ║")
			fmt.Println("║        menyenangkan dan hemat! 🌍💰   ║")
			fmt.Println("╚═══════════════════════════════════════╝")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func isiDataDummy(data *arrPengeluaran, count *int) {
	data[0] = Pengeluaran{"Transportasi", 500000}
	data[1] = Pengeluaran{"Hotel", 750000}
	data[2] = Pengeluaran{"Makanan", 300000}
	data[3] = Pengeluaran{"Hiburan", 200000}
	data[4] = Pengeluaran{"Oleh-oleh", 150000}
	*count = 5
}

func tampilkanMenu() {
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║         💸 MENU PENGELOLAAN BUDGET 💸          ║")
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

		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║       TAMBAH PENGELUARAN BARU      ║")
		fmt.Println("╚════════════════════════════════════╝")

		for pilihan < 1 || pilihan > 5 {
			fmt.Println(" Pilih kategori pengeluaran:")
			fmt.Println("1. Transportasi")
			fmt.Println("2. Hotel")
			fmt.Println("3. Makanan")
			fmt.Println("4. Hiburan")
			fmt.Println("5. Oleh-oleh")
			fmt.Print("📝 Pilihan (1-5): ")
			fmt.Scanln(&pilihan)

			if pilihan < 1 || pilihan > 5 {
				fmt.Println("❌ Pilihan tidak valid. Silakan coba lagi.")
			}
		}

		switch pilihan {
		case 1:
			kat = "Transportasi"
		case 2:
			kat = "Hotel"
		case 3:
			kat = "Makanan"
		case 4:
			kat = "Hiburan"
		case 5:
			kat = "Oleh-oleh"
		}

		fmt.Print("💰 Masukkan jumlah pengeluaran: ")
		fmt.Scanln(&jml)

		data[*count] = Pengeluaran{kat, jml}
		*count++
		fmt.Println("✅ Data berhasil ditambahkan!")
	} else {
		fmt.Println("⚠️ Data penuh. Tidak bisa menambah lagi.")
	}
}

func ubahData(data *arrPengeluaran, count int) {
	var index int
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║       Ubah Pengeluaran       ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Println("Daftar Pengeluaran Saat Ini:")
	fmt.Println("──────────────────────────────")
	for i := 0; i < count; i++ {
		fmt.Printf("│ %d. %-13s - Rp %d\n", i, data[i].kategori, data[i].jumlah)
	}
	fmt.Println("──────────────────────────────")
	fmt.Print("Pilih nomor pengeluaran yang ingin diubah: ")
	fmt.Scanln(&index)

	if index >= 0 && index < count {
		var pilihan int
		for pilihan < 1 || pilihan > 5 {
			fmt.Println("╔══════════════════════════════╗")
			fmt.Println("║      Pilih Kategori Baru     ║")
			fmt.Println("╚══════════════════════════════╝")
			fmt.Println("1. Transportasi")
			fmt.Println("2. Hotel")
			fmt.Println("3. Makanan")
			fmt.Println("4. Hiburan")
			fmt.Println("5. Oleh-oleh")
			fmt.Print("Pilihan (1-5): ")
			fmt.Scanln(&pilihan)

			if pilihan < 1 || pilihan > 5 {
				fmt.Println("⚠️  Pilihan tidak valid. Silakan coba lagi.")
			}
		}

		switch pilihan {
		case 1:
			data[index].kategori = "Transportasi"
		case 2:
			data[index].kategori = "Hotel"
		case 3:
			data[index].kategori = "Makanan"
		case 4:
			data[index].kategori = "Hiburan"
		case 5:
			data[index].kategori = "Oleh-oleh"
		}

		fmt.Print("💰 Masukkan jumlah baru: ")
		fmt.Scanln(&data[index].jumlah)

		fmt.Println("✅ Data berhasil diubah.")
	} else {
		fmt.Println("❌ Indeks tidak valid.")
	}
}

func hapusData(data *arrPengeluaran, count *int) {
	var index int
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║       Hapus Pengeluaran      ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Println("Daftar Pengeluaran Saat Ini:")
	fmt.Println("──────────────────────────────")
	for i := 0; i < *count; i++ {
		fmt.Printf("│ %d. %-13s - Rp %d\n", i, data[i].kategori, data[i].jumlah)
	}
	fmt.Println("──────────────────────────────")
	fmt.Print("Masukkan nomor pengeluaran yang ingin dihapus: ")
	fmt.Scanln(&index)

	if index >= 0 && index < *count {
		for i := index; i < *count-1; i++ {
			data[i] = data[i+1]
		}
		*count--
		fmt.Println("✅ Data berhasil dihapus.")
	} else {
		fmt.Println("❌ Indeks tidak valid.")
	}
}

func tampilkanData(data arrPengeluaran, count int) {
	if count == 0 {
		fmt.Println("⚠️  Belum ada data pengeluaran.")
	} else {
		fmt.Println("╔══════════════════════════════╗")
		fmt.Println("║       Daftar Pengeluaran     ║")
		fmt.Println("╚══════════════════════════════╝")
		for i := 0; i < count; i++ {
			fmt.Printf("│ %d. %-12s - Rp %d\n", i, data[i].kategori, data[i].jumlah)
		}
		fmt.Println("──────────────────────────────")
	}
}

func hitungTotalDanSaran(data arrPengeluaran, count int, budget int) {
	total := 0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}

	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Printf("║ 💰 Total pengeluaran: Rp %d\n", total)
	if total > budget {
		fmt.Printf("║ ⚠️  Anda melebihi budget sebesar: Rp -%d\n", total-budget)
		fmt.Println("║ 🚨 Kurangi pengeluaran Anda!")
	} else {
		fmt.Printf("║ ✅ Masih ada sisa budget: Rp %d\n", budget-total)
		fmt.Println("║ 🎉 Anda cukup hemat!")
	}
	fmt.Println("╚══════════════════════════════════════════════╝")
}

func cariDataKategori(data arrPengeluaran, count int) {
	var kat string
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║ 📂 Daftar Kategori:                    ║")
	fmt.Println("║ 1. Transportasi                        ║")
	fmt.Println("║ 2. Hotel                               ║")
	fmt.Println("║ 3. Makanan                             ║")
	fmt.Println("║ 4. Hiburan                             ║")
	fmt.Println("║ 5. Oleh-oleh                           ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Print("Masukkan nama kategori yang ingin dicari (Contoh: Makanan): ")
	fmt.Scanln(&kat)

	found := false
	fmt.Println("📋 Hasil Pencarian:")
	for i := 0; i < count; i++ {
		if data[i].kategori == kat {
			fmt.Printf("✅ Ditemukan di indeks %d: %s - Rp %d\n", i, data[i].kategori, data[i].jumlah)
			found = true
		}
	}

	if !found {
		fmt.Println("❌ Data tidak ditemukan!")
	}
}

func urutkanData(data *arrPengeluaran, count int) {
	var pilihan int
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║       📊 PILIH METODE PENGURUTAN       ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println("  1. Selection Sort berdasarkan jumlah   ")
	fmt.Println("  2. Insertion Sort berdasarkan kategori ")
	fmt.Print("Masukkan pilihan Anda (1-2): ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		var arah int
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║       URUT BERDASARKAN JUMLAH      ║")
		fmt.Println("╚════════════════════════════════════╝")
		fmt.Println("  1. Dari terkecil ke terbesar     ")
		fmt.Println("  2. Dari terbesar ke terkecil     ")
		fmt.Print("Masukkan pilihan arah (1-2): ")
		fmt.Scanln(&arah)

		if arah == 1 {
			for i := 0; i < count-1; i++ {
				min := i
				for j := i + 1; j < count; j++ {
					if data[j].jumlah < data[min].jumlah {
						min = j
					}
				}
				temp := data[i]
				data[i] = data[min]
				data[min] = temp
			}
			fmt.Println("\n✅ Data telah diurutkan berdasarkan jumlah (Terkecil ke Terbesar):")
		} else if arah == 2 {
			for i := 0; i < count-1; i++ {
				max := i
				for j := i + 1; j < count; j++ {
					if data[j].jumlah > data[max].jumlah {
						max = j
					}
				}
				temp := data[i]
				data[i] = data[max]
				data[max] = temp
			}
			fmt.Println("\n✅ Data telah diurutkan berdasarkan jumlah (Terbesar ke Terkecil):")
		} else {
			fmt.Println("❗ Pilihan arah pengurutan tidak valid.")
			return
		}

		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║         📋 HASIL PENGURUTAN        ║")
		fmt.Println("╚════════════════════════════════════╝")
		for i := 0; i < count; i++ {
			fmt.Printf(" %d. %-12s - %d\n", i, data[i].kategori, data[i].jumlah)
		}

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
		fmt.Println("\n✅ Data telah diurutkan berdasarkan kategori (A-Z):")
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║         📋 HASIL PENGURUTAN        ║")
		fmt.Println("╚════════════════════════════════════╝")
		for i := 0; i < count; i++ {
			fmt.Printf(" %d. %-12s - %d\n", i, data[i].kategori, data[i].jumlah)
		}
	} else {
		fmt.Println("❗ Pilihan tidak valid.")
	}
}

func tampilkanLaporan(data arrPengeluaran, count int, budget int) {
	if count == 0 {
		fmt.Println("╔════════════════════════════════════╗")
		fmt.Println("║     Tidak ada data pengeluaran.    ║")
		fmt.Println("╚════════════════════════════════════╝")
		return
	}

	fmt.Println("╔════════════════════════════════════╗")
	fmt.Println("║         LAPORAN PENGELUARAN        ║")
	fmt.Println("╚════════════════════════════════════╝")
	tampilkanData(data, count)

	total := 0
	for i := 0; i < count; i++ {
		total += data[i].jumlah
	}
	selisih := total - budget

	fmt.Println("╔════════════════════════════════════╗")
	fmt.Printf("║ Total Pengeluaran : %-14d ║\n", total)
	fmt.Printf("║ Budget            : %-14d ║\n", budget)

	// Menampilkan minus jika total pengeluaran melebihi budget
	if selisih < 0 {
		fmt.Printf("║ Selisih           : %-14d ║\n", -selisih) // untuk menampilkan plus
	} else {
		fmt.Printf("║ Selisih           : -%-13d ║\n", selisih) // untuk menampilkan minus
	}
	fmt.Println("╚════════════════════════════════════╝")
}
