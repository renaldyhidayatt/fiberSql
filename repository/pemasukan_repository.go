package repository

import (
	"database/sql"
	"fmt"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetAllPemasukan(db *sql.DB) ([]dto.Pemasukan, error) {
	pemasukan := []dto.Pemasukan{}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM pemasukan")

	if err != nil {
		return nil, fmt.Errorf("error, not connected to table, %w", err)
	}

	for rows.Next() {
		result := dto.Pemasukan{}

		if err := rows.Scan(&result.ID, &result.TGL_PEMASUKAN, &result.ID_SUMBER, &result.TGL_PEMASUKAN); err != nil {
			return nil, fmt.Errorf("Error, %w", err)
		}

		pemasukan = append(pemasukan, result)
	}

	return pemasukan, nil
}

func GetPemasukanID(id string, db *sql.DB) (dto.Pemasukan, error) {
	pemasukan := dto.Pemasukan{}

	rows, err := db.Query("SELECT * FROM pengeluaran WHERE id = $1", id)

	if err != nil {
		return pemasukan, nil
	}

	for rows.Next() {
		if err = rows.Scan(&pemasukan.ID, &pemasukan.JUMLAH, &pemasukan.TGL_PEMASUKAN, &pemasukan.ID_SUMBER); err != nil {
			return pemasukan, nil
		}
	}

	return pemasukan, nil
}

func CreatePemasukan(b *dto.Pemasukan, db *sql.DB) error {
	_, err := db.Query("INSERT INTO pemasukan (tgl_pemasukan, jumlah, id_sumber) VALUES ($1,$2,$3)", b.TGL_PEMASUKAN, b.JUMLAH, b.ID_SUMBER)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePemasukan(id string, b *dto.Pemasukan, db *sql.DB) error {
	_, err := db.Query("UPDATE pemasukan SET tgl_pemasukan=$2,jumlah=$3,id_sumber=4 WHERE id = $1", id, b.TGL_PEMASUKAN, b.JUMLAH, b.ID_SUMBER)

	if err != nil {
		return err
	}

	return nil
}

func DeletePemasukan(id string, db *sql.DB) error {
	_, err := db.Query("DELETE FROM pemasukan WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
