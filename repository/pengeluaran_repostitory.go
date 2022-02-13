package repository

import (
	"database/sql"
	"fmt"

	"github.com/renaldyhidayatt/pencobaan/dto"
)

func GetAllPengeluaran(db *sql.DB) ([]dto.Pengeluaran, error) {
	pengeluaran := []dto.Pengeluaran{}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM pengeluaran")

	if err != nil {
		return nil, fmt.Errorf("error, not connected to table, %w", err)
	}

	for rows.Next() {
		result := dto.Pengeluaran{}

		if err := rows.Scan(&result.ID, &result.TGL_PENGELUARAN, &result.ID_SUMBER, &result.TGL_PENGELUARAN); err != nil {
			return nil, fmt.Errorf("Error, %w", err)
		}

		pengeluaran = append(pengeluaran, result)
	}

	return pengeluaran, nil
}

func GetPengeluaranID(id string, db *sql.DB) (dto.Pengeluaran, error) {
	pengeluaran := dto.Pengeluaran{}

	rows, err := db.Query("SELECT * FROM pengeluaran WHERE id = $1", id)

	if err != nil {
		return pengeluaran, nil
	}

	for rows.Next() {
		if err = rows.Scan(&pengeluaran.ID, &pengeluaran.JUMLAH, &pengeluaran.TGL_PENGELUARAN, &pengeluaran.ID_SUMBER); err != nil {
			return pengeluaran, nil
		}
	}

	return pengeluaran, nil
}

func CreatePengeluaran(b *dto.Pengeluaran, db *sql.DB) error {
	_, err := db.Query("INSERT INTO pengeluaran (tgl_pengeluaran, jumlah, id_sumber) VALUES ($1,$2,$3)", b.TGL_PENGELUARAN, b.JUMLAH, b.ID_SUMBER)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePengeluaran(id string, b *dto.Pengeluaran, db *sql.DB) error {
	_, err := db.Query("UPDATE pengeluaran SET tgl_pengeluaran=$2,jumlah=$3,id_sumber=4 WHERE id = $1", id, b.TGL_PENGELUARAN, b.JUMLAH, b.ID_SUMBER)

	if err != nil {
		return err
	}

	return nil
}

func DeletePengeluaran(id string, db *sql.DB) error {
	_, err := db.Query("DELETE FROM pengeluaran WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
