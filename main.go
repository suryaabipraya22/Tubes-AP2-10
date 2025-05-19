package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Expense struct {
	Category string
	Amount   float64
	Note     string
}

var expenses []Expense
var plannedBudget float64

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Data dummy awal
	expenses = []Expense{
		{"Transportasi", 800000, "Tiket kereta Bandung-Jakarta"},
		{"Makanan", 150000, "Makan 3x sehari"},
		{"Akomodasi", 1200000, "Hotel 3 malam"},
		{"Oleh-oleh", 250000, "Belanja oleh-oleh"},
		{"Transportasi", 50000, "Ojek ke hotel"},
		{"Makanan", 100000, "Makan malam"},
		{"Wisata", 200000, "Tiket masuk tempat wisata"},
	}

	plannedBudget = 3000000

	for {
		fmt.Println("\n=== Menu Pengelola Budget Travelling ===")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Tampilkan Semua Pengeluaran")
		fmt.Println("3. Ubah Pengeluaran")
		fmt.Println("4. Hapus Pengeluaran")
		fmt.Println("5. Pencarian Pengeluaran (Kategori)")
		fmt.Println("6. Urutkan Pengeluaran")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addExpenseCLI(scanner)
		case "2":
			displayExpenses()
		case "3":
			updateExpenseCLI(scanner)
		case "4":
			deleteExpenseCLI(scanner)
		case "5":
			searchExpenseCLI(scanner)
		case "6":
			sortExpensesCLI(scanner)
		case "7":
			printReport()
		case "0":
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addExpenseCLI(scanner *bufio.Scanner) {
	fmt.Print("Kategori: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Print("Jumlah (Rp): ")
	scanner.Scan()
	amount, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Catatan: ")
	scanner.Scan()
	note := scanner.Text()

	expenses = append(expenses, Expense{category, amount, note})
	fmt.Println("✅ Pengeluaran berhasil ditambahkan.")
}

func updateExpenseCLI(scanner *bufio.Scanner) {
	displayExpenses()
	fmt.Print("Pilih nomor pengeluaran yang ingin diubah: ")
	scanner.Scan()
	index, _ := strconv.Atoi(scanner.Text())

	if index < 0 || index >= len(expenses) {
		fmt.Println("❌ Index tidak valid.")
		return
	}

	fmt.Print("Kategori baru: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Print("Jumlah baru (Rp): ")
	scanner.Scan()
	amount, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Catatan baru: ")
	scanner.Scan()
	note := scanner.Text()

	expenses[index] = Expense{category, amount, note}
	fmt.Println("✅ Data berhasil diubah.")
}

func deleteExpenseCLI(scanner *bufio.Scanner) {
	displayExpenses()
	fmt.Print("Pilih nomor pengeluaran yang ingin dihapus: ")
	scanner.Scan()
	index, _ := strconv.Atoi(scanner.Text())

	if index < 0 || index >= len(expenses) {
		fmt.Println("❌ Index tidak valid.")
		return
	}

	expenses = append(expenses[:index], expenses[index+1:]...)
	fmt.Println("✅ Data berhasil dihapus.")
}

func searchExpenseCLI(scanner *bufio.Scanner) {
	fmt.Print("Masukkan kategori untuk dicari: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Println("🔍 Hasil Sequential Search:")
	for i, e := range expenses {
		if strings.EqualFold(e.Category, category) {
			fmt.Printf("[%d] %s: Rp%.2f - %s\n", i, e.Category, e.Amount, e.Note)
		}
	}

	// Binary Search - harus diurutkan dulu berdasarkan kategori
	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].Category < expenses[j].Category
	})

	fmt.Println("🔎 Hasil Binary Search:")
	idx := binarySearchByCategory(expenses, category)
	if idx != -1 {
		e := expenses[idx]
		fmt.Printf("[%d] %s: Rp%.2f - %s\n", idx, e.Category, e.Amount, e.Note)
	} else {
		fmt.Println("Tidak ditemukan.")
	}
}

func sortExpensesCLI(scanner *bufio.Scanner) {
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. Jumlah (Selection Sort)")
	fmt.Println("2. Kategori (Insertion Sort)")
	fmt.Print("Pilih: ")

	scanner.Scan()
	pilihan := scanner.Text()

	switch pilihan {
	case "1":
		selectionSortByAmount(expenses)
		fmt.Println("✅ Diurutkan berdasarkan jumlah.")
	case "2":
		insertionSortByCategory(expenses)
		fmt.Println("✅ Diurutkan berdasarkan kategori.")
	default:
		fmt.Println("❌ Pilihan tidak valid.")
	}

	displayExpenses()
}

func printReport() {
	categoryMap := map[string]float64{}
	for _, e := range expenses {
		categoryMap[e.Category] += e.Amount
	}

	fmt.Println("\n📋 Laporan Pengeluaran per Kategori:")
	for cat, amt := range categoryMap {
		fmt.Printf("- %s: Rp%.2f\n", cat, amt)
	}

	total := totalExpense()
	fmt.Printf("\n🎯 Total Pengeluaran: Rp%.2f dari Anggaran Rp%.2f\n", total, plannedBudget)
	fmt.Printf("💰 Selisih: Rp%.2f\n", plannedBudget-total)

	suggestSaving()
}

func displayExpenses() {
	fmt.Println("\n📦 Daftar Pengeluaran:")
	for i, e := range expenses {
		fmt.Printf("[%d] %s - Rp%.2f (%s)\n", i, e.Category, e.Amount, e.Note)
	}
}

func totalExpense() float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}

func suggestSaving() {
	selisih := plannedBudget - totalExpense()
	if selisih < 0 {
		fmt.Printf("⚠ Anda melebihi anggaran sebesar Rp%.2f\n", -selisih)
	} else {
		fmt.Printf("✅ Anda masih hemat sebesar Rp%.2f\n", selisih)
	}
}

func selectionSortByAmount(exp []Expense) {
	for i := 0; i < len(exp)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(exp); j++ {
			if exp[j].Amount < exp[minIdx].Amount {
				minIdx = j
			}
		}
		exp[i], exp[minIdx] = exp[minIdx], exp[i]
	}
}

func insertionSortByCategory(exp []Expense) {
	for i := 1; i < len(exp); i++ {
		key := exp[i]
		j := i - 1
		for j >= 0 && exp[j].Category > key.Category {
			exp[j+1] = exp[j]
			j--
		}
		exp[j+1] = key
	}
}

func binarySearchByCategory(sorted []Expense, target string) int {
	low, high := 0, len(sorted)-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(sorted[mid].Category, target) {
			return mid
		} else if sorted[mid].Category < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
