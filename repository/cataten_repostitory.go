package repository

import (
	"database/sql"
	"fmt"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetAllCatatan(db *sql.DB) ([]dto.Catatan, error) {
	catatan := []dto.Catatan{}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM catatan")

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	for rows.Next() {
		result := dto.Catatan{}

		if err := rows.Scan(&result.ID, &result.CatatanText); err != nil {
			return nil, fmt.Errorf("error, not connected to database, %w", err)
		}

		catatan = append(catatan, result)
	}

	return catatan, nil

}

func GetCatatanID(id string, db *sql.DB) (dto.Catatan, error) {
	catatan := dto.Catatan{}

	rows, err := db.Query("SELECT * FROM catatan WHERE id=$1", id)

	if err != nil {
		return catatan, nil
	}

	for rows.Next() {
		if err = rows.Scan(&catatan.ID, &catatan.CatatanText); err != nil {
			return catatan, nil
		}
	}

	return catatan, nil
}

func CreateCatatan(b *dto.Catatan, db *sql.DB) error {

	_, err := db.Query("INSERT INTO catatan (catatan_text) VALUES ($1)", b.CatatanText)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCatatan(id int, b *dto.Catatan, db *sql.DB) error {
	_, err := db.Query("UPDATE catatan SET catatan_text = $2 WHERE id = $1", id, b.CatatanText)

	if err != nil {
		return err
	}

	return nil

}

func DeleteCatatan(id string, db *sql.DB) error {
	_, err := db.Query("DELETE FROM catatan WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
