package repository

import (
	"database/sql"
	"fmt"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetAllUang(db *sql.DB) ([]dto.Uang, error) {
	uang := []dto.Uang{}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM uang")

	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}

	for rows.Next() {
		result := dto.Uang{}

		if err := rows.Scan(&result.ID, &result.TGLUANG, &result.ID_PENGELUARAN, &result.ID_PENDAPATAN, &result.JUMLAH); err != nil {
			return nil, fmt.Errorf("Error %w", err)
		}

		uang = append(uang, result)
	}
	return uang, nil
}

func GetUangID(id string, db *sql.DB) (dto.Uang, error) {
	uang := dto.Uang{}

	rows, err := db.Query("SELECT * FROM uang WHERE id=$1", id)

	if err != nil {
		return uang, nil
	}

	for rows.Next() {
		if err = rows.Scan(&uang.ID, &uang.TGLUANG, &uang.ID_PENGELUARAN, &uang.ID_PENDAPATAN, &uang.JUMLAH); err != nil {
			return uang, nil
		}
	}

	return uang, nil
}

func CreateUang(b *dto.Uang, db *sql.DB) error {
	_, err := db.Query("INSERT INTO uang (tgl_uang,id_pengeluaran,id_pendapatan,jumlah) VALUES ($1,$2,$3,$4)", b.TGLUANG, b.ID_PENGELUARAN, b.ID_PENDAPATAN, b.JUMLAH)

	if err != nil {
		return err
	}

	return nil

}

func UpdateUang(id string, b *dto.Uang, db *sql.DB) error {
	_, err := db.Query("UPDATE uang SET tgl_uang = $2, id_pengeluaran = $3, id_pendapatan = $4, jumlah = $5 WHERE id=$1", id, b.TGLUANG, b.ID_PENGELUARAN, b.ID_PENDAPATAN, b.JUMLAH)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUang(id string, db *sql.DB) error {
	_, err := db.Query("DELETE FROM uang WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
