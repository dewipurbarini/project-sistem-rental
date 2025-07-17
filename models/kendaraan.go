package models

import (
	"database/sql"
)

type Kendaraan struct {
	ID           int
	Tipe         string
	Merk         string
	NomorPolisi  string
	HargaSewa    float64
	Status       string
}

// GetAllKendaraan mengambil semua data kendaraan
func GetAllKendaraan() ([]Kendaraan, error) {
	rows, err := DB.Query("SELECT id, tipe, merk, nomor_polisi, harga_sewa, status FROM kendaraan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Kendaraan
	for rows.Next() {
		var k Kendaraan
		if err := rows.Scan(&k.ID, &k.Tipe, &k.Merk, &k.NomorPolisi, &k.HargaSewa, &k.Status); err != nil {
			return nil, err
		}
		list = append(list, k)
	}
	// cek error setelah iterasi
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

// GetKendaraanByID mengambil satu kendaraan berdasarkan id
func GetKendaraanByID(id int) (*Kendaraan, error) {
	k := &Kendaraan{}
	err := DB.QueryRow("SELECT id, tipe, merk, nomor_polisi, harga_sewa, status FROM kendaraan WHERE id = ?", id).
		Scan(&k.ID, &k.Tipe, &k.Merk, &k.NomorPolisi, &k.HargaSewa, &k.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return k, nil
}

// CreateKendaraan menambahkan kendaraan baru
func CreateKendaraan(k *Kendaraan) error {
	_, err := DB.Exec(
		"INSERT INTO kendaraan (tipe, merk, nomor_polisi, harga_sewa, status) VALUES (?, ?, ?, ?, ?)",
		k.Tipe, k.Merk, k.NomorPolisi, k.HargaSewa, k.Status,
	)
	return err
}

// UpdateKendaraan memperbarui kendaraan
func UpdateKendaraan(k *Kendaraan) error {
	_, err := DB.Exec(
		"UPDATE kendaraan SET tipe=?, merk=?, nomor_polisi=?, harga_sewa=?, status=? WHERE id=?",
		k.Tipe, k.Merk, k.NomorPolisi, k.HargaSewa, k.Status, k.ID,
	)
	return err
}

// DeleteKendaraan menghapus kendaraan
func DeleteKendaraan(id int) error {
	_, err := DB.Exec("DELETE FROM kendaraan WHERE id=?", id)
	return err
}
