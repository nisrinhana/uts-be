package model

import "time"

type Alumni struct {
	ID         int       `json:"id"`
	UserID      int       `json:"user_id"`
	NIM        string    `json:"nim"`
	Nama       string    `json:"nama"`
	Jurusan    string    `json:"jurusan"`
	Angkatan   int       `json:"angkatan"`
	TahunLulus int       `json:"tahun_lulus"`
	Email      string    `json:"email"`
	NoTelepon  string    `json:"no_telepon"`
	Alamat     string    `json:"alamat"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Hanya untuk response status pekerjaan
type AlumniStatusKerjaResponse struct {
	ID             int       `json:"id"`
	Nama           string    `json:"nama"`
	Jurusan        string    `json:"jurusan"`
	Angkatan       int       `json:"angkatan"`
	BidangIndustri string    `json:"bidang_industri"`
	NamaPerusahaan string    `json:"nama_perusahaan"`
	PosisiJabatan  string    `json:"posisi_jabatan"`
	TanggalMulai   time.Time `json:"tanggal_mulai_kerja"`
	GajiRange      string    `json:"gaji_range"`
	Count          int       `json:"count"`
}
