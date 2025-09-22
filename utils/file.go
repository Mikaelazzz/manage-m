package utils

import (
	"os"
)

// FileExists memeriksa apakah file ada
func FileExists(filename string) bool {
	// Cek apakah file ada
	info, err := os.Stat(filename)
	// Jika error karena file tidak ada, kembalikan false
	if os.IsNotExist(err) {
		return false
	}
	// Kembalikan true jika file ada dan bukan direktori
	return !info.IsDir()
}
