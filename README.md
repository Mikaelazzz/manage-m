# Management Maba 

Aplikasi console untuk mengelola data mahasiswa baru yang dibuat menggunakan bahasa pemrograman Go. Aplikasi ini menyediakan fitur CRUD (Create, Read, Update, Delete) untuk data mahasiswa baru dengan penyimpanan data dalam format JSON.

## 🚀 Fitur

- **Tambah Maba**: Menambahkan data mahasiswa baru dengan nama, NIM, dan jurusan
- **Lihat Semua Maba**: Menampilkan seluruh daftar mahasiswa baru yang telah tersimpan
- **Edit Data Maba**: Mengubah informasi mahasiswa baru berdasarkan ID
- **Hapus Maba**: Menghapus data mahasiswa baru berdasarkan ID
- **Persistent Storage**: Data disimpan dalam file JSON (`maba.json`)
- **Auto-Increment ID**: ID mahasiswa baru dihasilkan secara otomatis
- **Timestamp Tracking**: Menyimpan waktu pembuatan dan pembaruan data

## 📁 Struktur Proyek

```
maba/
├── main.go             # Entry point aplikasi
├── go.mod              # Go module file
├── models/
│   └── maba.go         # Model data Maba
├── services/
│   └── maba.go         # Business logic dan data management
├── display/
│   └── menu.go         # User interface dan input handling
└── utils/
    └── file.go         # Utility functions untuk file operations
```


## 📊 Model Data

```go
type Maba struct {
    ID        int       `json:"id"`
    Nama      string    `json:"nama"`
    NIM       string    `json:"nim"`
    Jurusan   string    `json:"jurusan"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## 🛠️ Cara Menggunakan

### Prasyarat
- Go versi 1.25.1 atau lebih baru

### Instalasi dan Menjalankan

1. **Clone atau download proyek ini**
    ```
    git clone https://github.com/Mikaelazzz/manage-m.git
    ```

2. **Masuk ke direktori proyek**
   ```bash
   cd manage-m
   ```

3. **Bangun aplikasi**
    ```
    go build .
    ```

4. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

### Menu Aplikasi

Setelah aplikasi berjalan, Anda akan melihat menu utama:

```
== Management Maba ==

1. Tambah Maba
2. Lihat Semua Maba
3. Edit Data Maba
4. Hapus Maba
5. Exit

Pilih menu [1-5]:
```

### Panduan Penggunaan

#### 1. Tambah Maba
- Pilih menu 1
- Masukkan nama mahasiswa
- Masukkan NIM mahasiswa
- Masukkan jurusan mahasiswa
- Data akan disimpan otomatis dengan ID yang di-generate

#### 2. Lihat Semua Maba
- Pilih menu 2
- Semua data mahasiswa akan ditampilkan beserta informasi timestampnya

#### 3. Edit Data Maba
- Pilih menu 3
- Daftar mahasiswa akan ditampilkan
- Masukkan ID mahasiswa yang ingin diedit
- Masukkan data baru (nama, NIM, jurusan)
- Data akan diperbarui dan timestamp akan di-update

#### 4. Hapus Maba
- Pilih menu 4
- Daftar mahasiswa akan ditampilkan
- Masukkan ID mahasiswa yang ingin dihapus
- Data akan dihapus dari sistem

#### 5. Exit
- Pilih menu 5 untuk keluar dari aplikasi

## 💾 Penyimpanan Data

- Data disimpan dalam file `maba.json` di direktori yang sama dengan executable
- Format data menggunakan JSON dengan indentasi untuk kemudahan pembacaan
- File akan dibuat otomatis jika belum ada
- Data akan dimuat ulang setiap kali aplikasi dijalankan

## 🔧 Komponen Utama

### MabaService
Service utama yang mengelola:
- CRUD operations untuk data mahasiswa
- Auto-increment ID generation
- JSON file persistence
- Data validation dan error handling

### Display Functions
- `ShowMainMenu()`: Menampilkan menu utama
- `ShowMabas()`: Menampilkan daftar mahasiswa
- `GetTextInput()`: Input teks dari user
- `GetIntInput()`: Input angka dengan validasi
- `ShowSuccess()`, `ShowError()`, `ShowInfo()`: Pesan status

### Model Maba
- Struktur data mahasiswa dengan timestamp
- Constructor `NewMaba()` untuk membuat instance baru
- Method `Update()` untuk memperbarui data

## 🚨 Error Handling

Aplikasi menangani berbagai jenis error:
- ID mahasiswa tidak ditemukan
- Input tidak valid
- File I/O errors
- JSON parsing errors

## 📝 Contoh Data JSON

```json
[
  {
    "id": 1,
    "nama": "Mikaela",
    "nim": "12345678",
    "jurusan": "Ilmu Padi",
    "created_at": "2025-09-22T10:30:00Z",
    "updated_at": "2025-09-22T10:30:00Z"
  }
]
```

## 👨‍💻 Author

- **Nama**: Vincentius Johanes Lwie Jaya
- **Mata Kuliah**: Pemrograman Jaringan

---

**Note**: Pastikan Kalian memiliki Go versi lebih baru untuk menjalankan aplikasi ini.