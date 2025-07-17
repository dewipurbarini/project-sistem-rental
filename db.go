package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() error {
	var err error

	// Ganti user & password sesuai konfigurasi MySQL kamu
	// Jika tidak pakai password, cukup: root:@tcp...
	dsn := "root:@tcp(127.0.0.1:3306)/rental_kendaraan?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return DB.Ping()
}


