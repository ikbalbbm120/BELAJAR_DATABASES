package BELAJAR_DATABASES

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.open("mysql", "root:@tcp(localhost:3309)/belajar_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConss(100)
	db.SetConnMaxIdletime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

/**
 * DATABASE POOLING
 * sql.DB di golang sebenarnya bukanlah sebuah koneksi ke database melainkan sebuah pool
 * ke database, atau dikenal dengan istilah konsep database pooling
 * Di dalam sql.DB, golang melakukan managemen koneksi ke database secara otomatis, hal 
 * ini menjadikan kita tidak perlu melakukan managemen koneksi secara manual
 * Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal 
 * koneksi yang dibuat oleh golang, sehingga tidak membanjiri koneksi ke database, karena
 * biasanya ada batas maksimal koneksi yang ditangani oleh database yang kita gunakan
 * * Pengaturan Database Pooling
 * (DB) SetIdleConns(number) untuk pengaturan berapa jumlah koneksi minimal yang dibuat
 * (DB) SetOpenConns(number) untuk pengaturan berapa jumlah koneksi maksimal yang dibuat 
 * (DB) SetConnMaxIdleTime(duration) untuk pengaturan berapa lama koneksi yang sudah tidak 
 * digunakan akan dihapus
 * (DB) SetConnMaxLifetime(duration) untuk pengaturan berapa lama koneksi yang boleh digunakan
 */