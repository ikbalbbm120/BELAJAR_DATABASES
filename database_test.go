package BELAJAR_DATABASES

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)


func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3309)/belajar_database")
	if err != nil {
		panic(err)
	}
	defer db.close()
}

/**
 * INSTALL DATABASE DRIVER
 * go get -u github.com/go-sql-driver/mysql
 * 
 * MEMBUAT KONEKSI KE DATABASE
 * Untuk melakukan koneksi ke database kita bisa membuat objek sql.BD menggunakan 
 * function sql.Open(driver,dataSourceName)
 * Untuk menggunakan database MySQL kita bisa menggunakan driver "mysql"
 * Sedangkan untuk dataSourceName 
 * username:password@tcp(host:port)/database_name
 * objek sql.DB yang dihasilkan setelah berhasil terkoneksi berupa pointer
 * Jika objek sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan
 * function close()
 * 
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