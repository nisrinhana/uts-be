package repository

import (
	"fmt"
	"strconv"
	"tugas4go/app/model"
	"tugas4go/database"

)


func GetAllPekerjaan(isAdmin bool) ([]model.PekerjaanAlumni, error) {
	query := `
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
		       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at
		FROM pekerjaan_alumni
	`
	if !isAdmin {
		query += " WHERE deleted_at IS NULL"
	}
	query += " ORDER BY created_at DESC"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulai, &p.TanggalSelesai, &p.StatusPekerjaan,
			&p.Deskripsi, &p.CreatedAt, &p.UpdatedAt, 
		)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// func GetPekerjaanByID(id string, isAdmin bool) (model.PekerjaanAlumni, error) {
// 	var p model.PekerjaanAlumni

// 	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
// 		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
// 		       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at
// 		FROM pekerjaan_alumni
// 		WHERE id=$1  AND deleted_at IS NULL`, id
// 	// if !isAdmin {
// 	// 	query += " AND deleted_at IS NULL"
// 	// }

// 	err := database.DB.QueryRow(query, id).Scan(
// 		&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
// 		&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
// 		&p.TanggalMulai, &p.TanggalSelesai, &p.StatusPekerjaan,
// 		&p.Deskripsi, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt,
// 	)
// 	if err != nil {
// 		return p, err
// 	}
// 	return p, nil
// }

func GetPekerjaanByID(id string) (model.PekerjaanAlumni, error) {
	var p model.PekerjaanAlumni
	err := database.DB.QueryRow(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
		       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at
		FROM pekerjaan_alumni
		WHERE id=$1 AND deleted_at IS NULL
	`, id).Scan(
		&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
		&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
		&p.TanggalMulai, &p.TanggalSelesai, &p.StatusPekerjaan,
		&p.Deskripsi, &p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		return p, err
	}
	return p, nil
}

func GetPekerjaanByAlumniID(alumniID string, isAdmin bool) ([]model.PekerjaanAlumni, error) {
	query := `
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
		       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at
		FROM pekerjaan_alumni
		WHERE id=$1 
	`
	if !isAdmin {
		query += " AND deleted_at IS NULL"
	}
	query += " ORDER BY created_at DESC"

	rows, err := database.DB.Query(query, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulai, &p.TanggalSelesai, &p.StatusPekerjaan,
			&p.Deskripsi, &p.CreatedAt, &p.UpdatedAt, 
		)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func CreatePekerjaan(p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(`
		INSERT INTO pekerjaan_alumni 
		(alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, 
		 gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, 
		 deskripsi_pekerjaan, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10, NOW(), NOW())
	`,
		p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
		p.LokasiKerja, p.GajiRange, p.TanggalMulai, p.TanggalSelesai,
		p.StatusPekerjaan, p.Deskripsi,
	)
	return err
}

func UpdatePekerjaan(id string, p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni SET 
			alumni_id=$1, nama_perusahaan=$2, posisi_jabatan=$3, bidang_industri=$4, 
			lokasi_kerja=$5, gaji_range=$6, tanggal_mulai_kerja=$7, tanggal_selesai_kerja=$8, 
			status_pekerjaan=$9, deskripsi_pekerjaan=$10, updated_at=NOW() 
		WHERE id=$11
	`,
		p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
		p.LokasiKerja, p.GajiRange, p.TanggalMulai, p.TanggalSelesai,
		p.StatusPekerjaan, p.Deskripsi, id,
	)
	return err
}

func DeletePekerjaan(id string) error {
	_, err := database.DB.Exec("DELETE FROM pekerjaan_alumni WHERE id=$1", id)
	return err
}

func GetPekerjaanWithPagination(search, sortBy, order string, limit, offset int) ([]model.PekerjaanAlumni, error) {
	query := fmt.Sprintf(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortBy, order)

	rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulai, &p.TanggalSelesai, &p.StatusPekerjaan, &p.Deskripsi, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func CountPekerjaan(search string) (int, error) {
	var total int
	err := database.DB.QueryRow(`SELECT COUNT(*) FROM pekerjaan_alumni WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1`, "%"+search+"%").Scan(&total)
	return total, err
}

// SoftDelete
func SoftDeletePekerjaan(id string, isAdmin bool, userID int) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("ID tidak valid: %v", err)
	}

	query := `
		UPDATE pekerjaan_alumni 
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
		AND (alumni_id = (
			SELECT id FROM alumni WHERE user_id = $2
		) OR $3 = TRUE)
	`

	res, err := database.DB.Exec(query, intID, userID, isAdmin)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("tidak ada data yang diupdate, kemungkinan bukan milik user ini atau sudah dihapus")
	}

	return nil
}

//trash
func GetTrashedPekerjaan() ([]model.TrashPekerjaan, error) {
	query := `
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
		       status_pekerjaan, deskripsi_pekerjaan, deleted_at
		FROM pekerjaan_alumni
		WHERE deleted_at IS NOT NULL
		ORDER BY deleted_at DESC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.TrashPekerjaan
	for rows.Next() {
		var t model.TrashPekerjaan
		err := rows.Scan(
			&t.ID, &t.AlumniID, &t.NamaPerusahaan, &t.PosisiJabatan,
			&t.BidangIndustri, &t.LokasiKerja, &t.GajiRange,
			&t.TanggalMulai, &t.TanggalSelesai, &t.StatusPekerjaan,
			&t.Deskripsi, &t.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}


// Restore data dari Trash
func RestorePekerjaan(id string) error {
	query := `
		UPDATE pekerjaan_alumni 
		SET deleted_at = NULL 
		WHERE id = $1 AND deleted_at IS NOT NULL
	`
	res, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("data tidak ditemukan atau sudah aktif")
	}
	return nil
}

//harddelete
func HardDeletePekerjaan(id string) error {
	query := `
		DELETE FROM pekerjaan_alumni 
		WHERE id = $1 AND deleted_at IS NOT NULL
	`
	res, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return fmt.Errorf("data tidak ditemukan atau belum dihapus (belum di-trash)")
	}

	return nil
}



