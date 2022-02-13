package repository

import (
	"database/sql"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetUserEmail(e string, db *sql.DB) (dto.Users, error) {
	users := dto.Users{}

	rows, err := db.Query("SELECT * FROM users WHERE email=$1", e)

	if err != nil {
		return users, nil
	}

	for rows.Next() {
		if err = rows.Scan(&users.ID, &users.Email, &users.Name, &users.Password); err != nil {
			return users, nil
		}
	}

	return users, nil
}

func CreateUsers(b *dto.Users, db *sql.DB) error {
	_, err := db.Query("INSERT INTO users (name,email,password) VALUES ($1,$2,$3)", b.Name, b.Email, b.Password)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUsers(id int, b *dto.Users, db *sql.DB) error {
	_, err := db.Query("UPDATE users SET name=$2,email=$3,password=$4 WHERE id = $1", id, b.Name, b.Email, b.Password)

	if err != nil {
		return err
	}

	return nil

}

func DeleteUsers(id int, db *sql.DB) error {
	_, err := db.Query("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
