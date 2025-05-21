package main

import "fmt"

const NMAX = 100

var data [NMAX]string
var count int

func main() {
	var pilihan int

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
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("==== MENU DATA ====")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Tampilkan Data")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahData() {
	if count < NMAX {
		var input string
		fmt.Print("Masukkan data baru: ")
		fmt.Scanln(&input)
		data[count] = input
		count++
		fmt.Println("Data berhasil ditambahkan.")
	} else {
		fmt.Println("Array penuh.")
	}
}

func ubahData() {
	var index int
	fmt.Print("Masukkan indeks data yang ingin diubah: ")
	fmt.Scanln(&index)
	if index >= 0 && index < count {
		var baru string
		fmt.Print("Masukkan data baru: ")
		fmt.Scanln(&baru)
		data[index] = baru
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusData() {
	var index int
	fmt.Print("Masukkan indeks data yang ingin dihapus: ")
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

func tampilkanData() {
	if count == 0 {
		fmt.Println("Belum ada data.")
	} else {
		fmt.Println("Isi Data:")
		for i := 0; i < count; i++ {
			fmt.Printf("%d. %s\n", i, data[i])
		}
	}
}
