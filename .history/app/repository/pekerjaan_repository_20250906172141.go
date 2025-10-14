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

// CREATE
func CreatePekerjaan(p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(
		`INSERT INTO pekerjaan_alumni 
        (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, 
         gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, 
         deskripsi_pekerjaan, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10, NOW(), NOW())`,
		p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
		p.GajiRange, p.TanggalMulai, p.TanggalSelesai, p.StatusPekerjaan, p.Deskripsi,
	)
	return err
}

// UPDATE
func UpdatePekerjaan(id string, p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(
		`UPDATE pekerjaan_alumni SET 
        alumni_id=$1, nama_perusahaan=$2, posisi_jabatan=$3, bidang_industri=$4, 
        lokasi_kerja=$5, gaji_range=$6, tanggal_mulai_kerja=$7, tanggal_selesai_kerja=$8, 
        status_pekerjaan=$9, deskripsi_pekerjaan=$10, updated_at=NOW() 
        WHERE id=$11`,
		p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
		p.GajiRange, p.TanggalMulai, p.TanggalSelesai, p.StatusPekerjaan, p.Deskripsi, id,
	)
	return err
}

// DELETE
func DeletePekerjaan(id string) error {
	_, err := database.DB.Exec("DELETE FROM pekerjaan_alumni WHERE id=$1", id)
	return err
}
