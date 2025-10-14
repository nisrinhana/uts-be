package repository

import (
	"fmt"                    
	"tugas4go/app/model"
	"tugas4go/database"
)

func GetAllAlumni() ([]model.Alumni, error) {
	rows, err := database.DB.Query("SELECT * FROM alumni ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Alumni
	for rows.Next() {
		var a model.Alumni
		err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
			&a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func GetAlumniByID(id string) (model.Alumni, error) {
	var a model.Alumni
	err := database.DB.QueryRow(
		"SELECT * FROM alumni WHERE id = $1", id,
	).Scan(
		&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus,
		&a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		return a, err
	}
	return a, nil
}

func CreateAlumni(a model.Alumni) error {
	_, err := database.DB.Exec(
		`INSERT INTO alumni 
        (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at) 
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8, NOW(), NOW())`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat,
	)
	return err
}

func UpdateAlumni(id string, a model.Alumni) error {
	_, err := database.DB.Exec(
		`UPDATE alumni SET nim=$1, nama=$2, jurusan=$3, angkatan=$4, tahun_lulus=$5, 
        email=$6, no_telepon=$7, alamat=$8, updated_at=NOW() WHERE id=$9`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus,
		a.Email, a.NoTelepon, a.Alamat, id,
	)
	return err
}

func DeleteAlumni(id string) error {
	_, err := database.DB.Exec("DELETE FROM alumni WHERE id=$1", id)
	return err
}
func GetAlumniWithPagination(search, sortBy, order string, limit, offset int) ([]model.Alumni, error) {
    query := fmt.Sprintf(`
        SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
        FROM alumni
        WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
        ORDER BY %s %s
        LIMIT $2 OFFSET $3
    `, sortBy, order)

    rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Alumni
    for rows.Next() {
        var a model.Alumni
        if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
            &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat,
            &a.CreatedAt, &a.UpdatedAt); err != nil {
            return nil, err
        }
        list = append(list, a)
    }
    return list, nil
}

func CountAlumni(search string) (int, error) {
    var total int
    err := database.DB.QueryRow(`
        SELECT COUNT(*) FROM alumni 
        WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1
    `, "%"+search+"%").Scan(&total)
    return total, err
}


// // //
func GetAlumniStatusKerjaLebih1Tahun() ([]model.AlumniStatusKerjaResponse, error) {
	query := `
		SELECT 
			a.id,
			a.nama,
			a.jurusan,
			a.angkatan,
			p.bidang_industri,
			p.nama_perusahaan,
			p.posisi_jabatan,
			p.tanggal_mulai_kerja,
			p.gaji_range,
			COUNT(p.id) AS count
		FROM alumni a
		JOIN pekerjaan_alumni p ON a.id = p.alumni_id
		WHERE AGE(COALESCE(p.tanggal_selesai_kerja, CURRENT_DATE), p.tanggal_mulai_kerja) > INTERVAL '1 year'
		GROUP BY a.id, a.nama, a.jurusan, a.angkatan, 
		         p.bidang_industri, p.nama_perusahaan, 
		         p.posisi_jabatan, p.tanggal_mulai_kerja, p.gaji_range
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.AlumniStatusKerjaResponse
	for rows.Next() {
		var a model.AlumniStatusKerjaResponse
		err := rows.Scan(
			&a.ID, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.BidangIndustri, &a.NamaPerusahaan,
			&a.PosisiJabatan, &a.TanggalMulai,
			&a.GajiRange, &a.Count,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil
}
