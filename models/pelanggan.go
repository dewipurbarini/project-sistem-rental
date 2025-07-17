package models

type Pelanggan struct {
	ID      int
	Nama    string
	Alamat  string
	Telepon string
}

func GetAllPelanggan() ([]Pelanggan, error) {
	rows, err := DB.Query("SELECT id, nama, alamat, telepon FROM pelanggan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Pelanggan
	for rows.Next() {
		var p Pelanggan
		if err := rows.Scan(&p.ID, &p.Nama, &p.Alamat, &p.Telepon); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

func GetPelangganByID(id int) (*Pelanggan, error) {
	p := &Pelanggan{}
	err := DB.QueryRow("SELECT id, nama, alamat, telepon FROM pelanggan WHERE id=?", id).
		Scan(&p.ID, &p.Nama, &p.Alamat, &p.Telepon)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func CreatePelanggan(p *Pelanggan) error {
	_, err := DB.Exec("INSERT INTO pelanggan (nama, alamat, telepon) VALUES (?,?,?)",
		p.Nama, p.Alamat, p.Telepon)
	return err
}

func UpdatePelanggan(p *Pelanggan) error {
	_, err := DB.Exec("UPDATE pelanggan SET nama=?, alamat=?, telepon=? WHERE id=?",
		p.Nama, p.Alamat, p.Telepon, p.ID)
	return err
}

func DeletePelanggan(id int) error {
	_, err := DB.Exec("DELETE FROM pelanggan WHERE id=?", id)
	return err
}
