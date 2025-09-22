package main

import (
	"bufio"
	"fmt"
	"hw1/display"
	"hw1/services"
	"os"
)

func main() {
	// Inisialisasi reader
	reader := bufio.NewReader(os.Stdin)
	// Inisialisasi service
	mabaService := services.NewMabaService("maba.json")

	for {
		// Tampilkan menu utama
		display.ShowMainMenu()
		// Dapatkan pilihan menu dari pengguna
		choice := display.GetIntInput("", reader)

		// Proses pilihan menu
		switch choice {
		case 1: // Tambah Maba
			nama := display.GetTextInput("Masukkan nama: ", reader)
			nim := display.GetTextInput("Masukkan NIM: ", reader)
			jurusan := display.GetTextInput("Masukkan jurusan: ", reader)

			// Tambah data maba
			mabaService.AddMaba(nama, nim, jurusan)
			display.ShowSuccess("[ SUCCESS ] Data maba berhasil ditambahkan!")

		case 2: // Lihat Semua Maba
			mabas := mabaService.GetMabas()
			// Tampilkan daftar maba
			display.ShowMabas(mabas, "[ INFO ] DATA MAHASISWA BARU")

		case 3: // Edit Data Maba
			display.ShowMabas(mabaService.GetMabas(), "[ INFO ] EDIT DATA MABA")
			// Cek apakah ada data maba
			if len(mabaService.GetMabas()) == 0 {
				break
			}
			id := display.GetIntInput("[ INFO ] Masukkan ID maba yang akan diedit : ", reader)
			// Dapatkan data maba berdasarkan ID
			maba, err := mabaService.GetMabaByID(id)
			if err != nil {
				display.ShowError(err.Error())
				break
			}

			fmt.Println("\n[ INFO ]")
			fmt.Printf("Nama : %s\n", maba.Nama)
			fmt.Printf("NIM : %s\n", maba.NIM)
			fmt.Printf("Jurusan : %s\n", maba.Jurusan)

			nama := display.GetTextInput("[ INFO ] Masukkan nama baru : ", reader)
			nim := display.GetTextInput("[ INFO ] Masukkan NIM baru : ", reader)
			jurusan := display.GetTextInput("[ INFO ] Masukkan jurusan baru : ", reader)

			// Perbarui data maba
			err = mabaService.UpdateMaba(id, nama, nim, jurusan)
			if err != nil {
				// Tampilkan pesan error jika ada
				display.ShowError(err.Error())
			} else {
				display.ShowSuccess("[ SUCCESS ] Data maba berhasil diperbarui!")
			}

		case 4: // Hapus Maba
			display.ShowMabas(mabaService.GetMabas(), "[ INFO ] HAPUS DATA MABA")
			// Cek apakah ada data maba
			if len(mabaService.GetMabas()) == 0 {
				break
			}
			id := display.GetIntInput("[ INFO ] Masukkan ID maba yang akan dihapus : ", reader)
			// Hapus data maba berdasarkan ID
			err := mabaService.DeleteMaba(id)
			if err != nil {
				display.ShowError(err.Error())
			} else {
				display.ShowSuccess("[ SUCCESS ] Data maba berhasil dihapus!")
			}

		case 5: // Exit
			display.ShowInfo("[ INFO ] Terima kasih!")
			return

		default:
			display.ShowError("[ ERROR ] Pilihan menu tidak valid !!")
		}
	}
}
