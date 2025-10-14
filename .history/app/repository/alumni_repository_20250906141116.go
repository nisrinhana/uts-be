package repository

import (
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
