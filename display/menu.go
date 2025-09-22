package display

import (
	"bufio"
	"fmt"
	"hw1/models"
	"strconv"
	"strings"
)

// Menampilkan menu utama
func ShowMainMenu() {
	fmt.Println(" == Management Maba ==")
	fmt.Println("")
	fmt.Println("1. Tambah Maba")
	fmt.Println("2. Lihat Semua Maba")
	fmt.Println("3. Edit Data Maba")
	fmt.Println("4. Hapus Maba")
	fmt.Println("5. Exit")
	fmt.Println("")
	fmt.Print("Pilih menu [1-5]: ")
}

// Mendapatkan input teks dari pengguna
func GetTextInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Mendapatkan input angka dari pengguna dengan validasi
func GetIntInput(prompt string, reader *bufio.Reader) int {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.TrimSpace(text))
		if err == nil {
			return num
		}
		fmt.Println("Input tidak valid! Masukkan angka.")
	}
}

// Menampilkan daftar mahasiswa baru
func ShowMabas(mabas []*models.Maba, title string) {
	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Printf(" %s\n", title)
	fmt.Println(strings.Repeat("-", 50))

	// Cek apakah ada data maba
	if len(mabas) == 0 {
		fmt.Println("Tidak ada data maba!")
		return
	}

	// Menampilkan data maba
	for _, maba := range mabas {
		fmt.Printf("ID: %d\n", maba.ID)
		fmt.Printf("Nama: %s\n", maba.Nama)
		fmt.Printf("NIM: %s\n", maba.NIM)
		fmt.Printf("Jurusan: %s\n", maba.Jurusan)
		// Menampilkan timestamp dengan format yang lebih baik
		fmt.Printf("Dibuat: %s\n", maba.CreatedAt.Format("2006-01-02 15:04:05"))
		// Hanya tampilkan updatedAt jika berbeda dengan createdAt
		if maba.UpdatedAt.After(maba.CreatedAt) {
			fmt.Printf("Diperbarui: %s\n", maba.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	}
}

// Menampilkan pesan sukses
func ShowSuccess(message string) {
	fmt.Printf("\n[ SUCCESS ] %s\n", message)
}

// Menampilkan pesan error
func ShowError(message string) {
	fmt.Printf("\n[ ERROR ] %s\n", message)
}

// Menampilkan pesan informasi
func ShowInfo(message string) {
	fmt.Printf("\n[ INFO ] %s\n", message)
}
