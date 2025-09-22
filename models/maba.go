package models

import (
	"time"
)

// Maba merepresentasikan data mahasiswa
type Maba struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	NIM       string    `json:"nim"`
	Jurusan   string    `json:"jurusan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewMaba adalah konstruktor untuk membuat instance Maba
func NewMaba(id int, nama, nim, jurusan string) *Maba {
	now := time.Now()
	return &Maba{
		ID:        id,
		Nama:      nama,
		NIM:       nim,
		Jurusan:   jurusan,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Update memperbarui data Maba dan mengatur waktu pembaruan
func (m *Maba) Update(nama, nim, jurusan string) {
	m.Nama = nama
	m.NIM = nim
	m.Jurusan = jurusan
	m.UpdatedAt = time.Now()
}
