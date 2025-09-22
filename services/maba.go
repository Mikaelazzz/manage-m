package services

import (
	"encoding/json"
	"errors"
	"hw1/models"
	"os"
)

// MabaService mengelola data mahasiswa baru
type MabaService struct {
	// Menyimpan data mahasiswa baru
	mabas   []*models.Maba
	nextID  int
	storage string
}

// NewMabaService membuat instance baru dari MabaService
func NewMabaService(storageFile string) *MabaService {
	// Inisialisasi dengan data dari file jika ada
	service := &MabaService{
		// Inisialisasi slice
		mabas: make([]*models.Maba, 0),
		// Inisialisasi ID berikutnya
		nextID: 1,
		// Inisialisasi storage
		storage: storageFile,
	}
	// Muat data dari file
	service.loadFromFile()
	return service
}

// Menambahkan mahasiswa baru
func (s *MabaService) AddMaba(nama, nim, jurusan string) *models.Maba {
	// Validasi input
	maba := models.NewMaba(s.nextID, nama, nim, jurusan)
	// Tambahkan ke slice
	s.mabas = append(s.mabas, maba)
	s.nextID++
	// Simpan ke file setelah penambahan
	s.saveToFile()
	return maba
}

// Mengambil semua mahasiswa baru
func (s *MabaService) GetMabas() []*models.Maba {
	return s.mabas
}

// Mengambil mahasiswa baru berdasarkan ID
func (s *MabaService) GetMabaByID(id int) (*models.Maba, error) {
	// Cari maba berdasarkan ID
	for _, maba := range s.mabas {
		if maba.ID == id {
			return maba, nil
		}
	}
	return nil, errors.New("ID Maba tidak ditemukan")
}

// Mengupdate data mahasiswa baru berdasarkan ID
// Mengembalikan error jika ID tidak ditemukan
// atau input tidak valid
func (s *MabaService) UpdateMaba(id int, nama, nim, jurusan string) error {
	// Validasi input
	maba, err := s.GetMabaByID(id)
	if err != nil {
		return err
	}
	// Update data
	maba.Update(nama, nim, jurusan)
	// Simpan ke file setelah update
	s.saveToFile()
	return nil
}

// Menghapus mahasiswa baru berdasarkan ID
// Mengembalikan error jika ID tidak ditemukan
func (s *MabaService) DeleteMaba(id int) error {
	// Cari dan hapus maba berdasarkan ID
	for i, maba := range s.mabas {
		// Jika ditemukan, hapus dari slice
		if maba.ID == id {
			// Hapus elemen dari slice
			s.mabas = append(s.mabas[:i], s.mabas[i+1:]...)
			// Simpan ke file setelah penghapusan
			s.saveToFile()
			return nil
		}
	}
	return errors.New("ID Maba tidak ditemukan")
}

// Menyimpan data ke file JSON
func (s *MabaService) saveToFile() error {
	// Marshal data ke JSON
	data, err := json.MarshalIndent(s.mabas, "", "  ")
	if err != nil {
		return err
	}
	// Tulis ke file
	return os.WriteFile(s.storage, data, 0644)
}

// Memuat data dari file JSON
func (s *MabaService) loadFromFile() error {
	// Baca file
	data, err := os.ReadFile(s.storage)
	if err != nil {
		// Jika terjadi error, periksa apakah file tidak ada
		if os.IsNotExist(err) {
			return nil // File tidak ada, mulai dengan data kosong
		}
		return err
	}

	// Parse JSON
	var mabas []*models.Maba
	// Unmarshal data JSON ke slice
	err = json.Unmarshal(data, &mabas)
	if err != nil {
		return err
	}

	// Set data dan nextID
	s.mabas = mabas
	// Tentukan nextID berdasarkan ID tertinggi yang ada
	if len(mabas) > 0 {
		// Asumsi data sudah terurut berdasarkan ID
		s.nextID = mabas[len(mabas)-1].ID + 1
	}
	return nil
}
