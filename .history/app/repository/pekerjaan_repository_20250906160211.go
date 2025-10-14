package repository

import (
	"tugas4go/app/model"
	"tugas4go/database"
)

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
