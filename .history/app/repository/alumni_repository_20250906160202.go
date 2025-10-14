package repository

import (
	"tugas4go/app/model"
	"tugas4go/database"
)

// GET semua
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

// GET by ID
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

// CREATE alumni
func CreateAlumni(a model.Alumni) error {
	_, err := database.DB.Exec(
		`INSERT INTO alumni 
        (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at) 
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8, NOW(), NOW())`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat,
	)
	return err
}

// UPDATE alumni
func UpdateAlumni(id string, a model.Alumni) error {
	_, err := database.DB.Exec(
		`UPDATE alumni SET nim=$1, nama=$2, jurusan=$3, angkatan=$4, tahun_lulus=$5, 
        email=$6, no_telepon=$7, alamat=$8, updated_at=NOW() WHERE id=$9`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus,
		a.Email, a.NoTelepon, a.Alamat, id,
	)
	return err
}

// DELETE alumni
func DeleteAlumni(id string) error {
	_, err := database.DB.Exec("DELETE FROM alumni WHERE id=$1", id)
	return err
}
