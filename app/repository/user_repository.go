package repository

import (
	"tugas4go/app/model"
	"tugas4go/database"
)

func GetUserByUsername(username string) (model.User, error) {
	var u model.User
	err := database.DB.QueryRow(`
		SELECT id, username, email, password_hash, role, created_at, updated_at
		FROM users
		WHERE username = $1
	`, username).Scan(
		&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt, &u.UpdatedAt,
	)
	return u, err
}


func GetUserByEmail(email string) (model.User, error) {
	var u model.User
	err := database.DB.QueryRow(`
		SELECT id, username, email, password_hash, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`, email).Scan(
		&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.Role, &u.CreatedAt, &u.UpdatedAt,
	)
	return u, err
}
