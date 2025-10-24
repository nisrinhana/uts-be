package model

import (
	"database/sql"
	"time"
)

// Struct utama untuk data pekerjaan alumni (data aktif)
type PekerjaanAlumni struct {
	ID              int            `json:"id"`
	AlumniID        int            `json:"alumni_id"`
	NamaPerusahaan  string         `json:"nama_perusahaan"`
	PosisiJabatan   string         `json:"posisi_jabatan"`
	BidangIndustri  string         `json:"bidang_industri"`
	LokasiKerja     string         `json:"lokasi_kerja"`
	GajiRange       string         `json:"gaji_range"`
	TanggalMulai    time.Time      `json:"tanggal_mulai_kerja"`
	TanggalSelesai  *time.Time     `json:"tanggal_selesai_kerja"`
	StatusPekerjaan string         `json:"status_pekerjaan"`
	Deskripsi       string         `json:"deskripsi_pekerjaan"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	// DeletedAt       *sql.NullTime  `json:"deleted_at"`
}

// (Trash)
type TrashPekerjaan struct {
	ID              int            `json:"id"`
	AlumniID        int            `json:"alumni_id"`
	NamaPerusahaan  string         `json:"nama_perusahaan"`
	PosisiJabatan   string         `json:"posisi_jabatan"`
	BidangIndustri  string         `json:"bidang_industri"`
	LokasiKerja     string         `json:"lokasi_kerja"`
	GajiRange       string         `json:"gaji_range"`
	TanggalMulai    time.Time      `json:"tanggal_mulai_kerja"`
	TanggalSelesai  *time.Time     `json:"tanggal_selesai_kerja"`
	StatusPekerjaan string         `json:"status_pekerjaan"`
	Deskripsi       string         `json:"deskripsi_pekerjaan"`
	DeletedAt       sql.NullTime   `json:"deleted_at"` 
}
