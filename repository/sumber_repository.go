package repository

import (
	"database/sql"
	"fmt"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetAllSumber(db *sql.DB) ([]dto.Sumber, error) {
	sumber := []dto.Sumber{}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM sumber")

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	for rows.Next() {
		result := dto.Sumber{}

		if err := rows.Scan(&result.ID, &result.Nama); err != nil {
			return nil, fmt.Errorf("error, not connected to database, %w", err)
		}

		sumber = append(sumber, result)
	}

	return sumber, nil

}

func GetSumberID(id string, db *sql.DB) (dto.Sumber, error) {
	sumber := dto.Sumber{}

	rows, err := db.Query("SELECT * FROM sumber WHERE id=$1", id)

	if err != nil {
		return sumber, nil
	}

	for rows.Next() {
		if err = rows.Scan(&sumber.ID, &sumber.Nama); err != nil {
			return sumber, nil
		}
	}

	return sumber, nil
}

func CreateSumber(b *dto.Sumber, db *sql.DB) error {

	_, err := db.Query("INSERT INTO sumber (nama) VALUES ($1)", b.Nama)

	if err != nil {
		return err
	}

	return nil
}

func UpdateSumber(id string, b *dto.Sumber, db *sql.DB) error {
	_, err := db.Query("UPDATE sumber SET nama = $2 WHERE id = $1", id, b.Nama)

	if err != nil {
		return err
	}

	return nil

}

func DeleteSumber(id string, db *sql.DB) error {
	_, err := db.Query("DELETE FROM sumber WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
