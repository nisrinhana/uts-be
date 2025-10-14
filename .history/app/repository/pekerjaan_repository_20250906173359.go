package repository

import (
	"tugas4go/app/model"
	"tugas4go/database"
)

// GET semua
func GetAllPekerjaan() ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query("SELECT * FROM pekerjaan_alumni ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulai,
			&p.TanggalSelesai, &p.StatusPekerjaan, &p.Deskripsi,
			&p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// GET by ID
func GetPekerjaanByID(id string) (model.PekerjaanAlumni, error) {
	var p model.PekerjaanAlumni
	err := database.DB.QueryRow("SELECT * FROM pekerjaan_alumni WHERE id=$1", id).
		Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulai,
			&p.TanggalSelesai, &p.StatusPekerjaan, &p.Deskripsi,
			&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p, err
	}
	return p, nil
}

// GET by alumni_id
func GetPekerjaanByAlumniID(alumniID string) ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query("SELECT * FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY created_at DESC", alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulai,
			&p.TanggalSelesai, &p.StatusPekerjaan, &p.Deskripsi,
			&p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// CREATE
func CreatePekerjaan(p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(
		`INSERT INTO pekerjaan_alumni 
        (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, 
         gaji_range, tanggal_mulai_kerja, tanggal_selesai_k_
