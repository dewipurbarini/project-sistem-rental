package models

import "database/sql"

type Transaksi struct {
	ID             int
	IDPelanggan    int
	IDKendaraan    int
	TanggalPinjam  string
	TanggalKembali string
	Total          float64
	Status         string
	PelangganNama  string
	KendaraanMerk  string
}

func GetAllTransaksi() ([]Transaksi, error) {
	rows, err := DB.Query(`
		SELECT t.id, t.pelanggan_id, p.nama, t.kendaraan_id, k.merk,
		       t.tanggal_pinjam, t.tanggal_kembali, t.total, t.status
		FROM transaksi t
		JOIN pelanggan p ON t.pelanggan_id = p.id
		JOIN kendaraan k ON t.kendaraan_id = k.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Transaksi
	for rows.Next() {
		var t Transaksi
		if err := rows.Scan(&t.ID, &t.IDPelanggan, &t.PelangganNama, &t.IDKendaraan, &t.KendaraanMerk,
			&t.TanggalPinjam, &t.TanggalKembali, &t.Total, &t.Status); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}

func GetTransaksiByID(id int) (*Transaksi, error) {
	var t Transaksi
	err := DB.QueryRow(`
		SELECT id, pelanggan_id, kendaraan_id, tanggal_pinjam, tanggal_kembali, total, status
		FROM transaksi WHERE id=?`, id).Scan(
		&t.ID, &t.IDPelanggan, &t.IDKendaraan, &t.TanggalPinjam, &t.TanggalKembali, &t.Total, &t.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func CreateTransaksi(t *Transaksi) error {
	_, err := DB.Exec(`
		INSERT INTO transaksi (pelanggan_id, kendaraan_id, tanggal_pinjam, tanggal_kembali, total, status)
		VALUES (?,?,?,?,?,?)`,
		t.IDPelanggan, t.IDKendaraan, t.TanggalPinjam, t.TanggalKembali, t.Total, t.Status,
	)
	return err
}

func UpdateTransaksi(t *Transaksi) error {
	_, err := DB.Exec(`
		UPDATE transaksi SET pelanggan_id=?, kendaraan_id=?, tanggal_pinjam=?, tanggal_kembali=?, total=?, status=?
		WHERE id=?`,
		t.IDPelanggan, t.IDKendaraan, t.TanggalPinjam, t.TanggalKembali, t.Total, t.Status, t.ID,
	)
	return err
}

func DeleteTransaksi(id int) error {
	_, err := DB.Exec(`DELETE FROM transaksi WHERE id=?`, id)
	return err
}

// Laporan bulanan
func GetTransaksiBulanan(month int, year int) ([]Transaksi, error) {
	rows, err := DB.Query(`
		SELECT id, pelanggan_id, kendaraan_id, tanggal_pinjam, tanggal_kembali, total, status
		FROM transaksi
		WHERE MONTH(tanggal_pinjam)=? AND YEAR(tanggal_pinjam)=?`, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Transaksi
	for rows.Next() {
		var t Transaksi
		if err := rows.Scan(&t.ID, &t.IDPelanggan, &t.IDKendaraan, &t.TanggalPinjam, &t.TanggalKembali, &t.Total, &t.Status); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}

// Laporan tahunan
func GetTransaksiYearly(year int) ([]Transaksi, error) {
	rows, err := DB.Query(`
		SELECT id, pelanggan_id, kendaraan_id, tanggal_pinjam, tanggal_kembali, total, status
		FROM transaksi
		WHERE YEAR(tanggal_pinjam)=?`, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Transaksi
	for rows.Next() {
		var t Transaksi
		if err := rows.Scan(&t.ID, &t.IDPelanggan, &t.IDKendaraan, &t.TanggalPinjam, &t.TanggalKembali, &t.Total, &t.Status); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}
