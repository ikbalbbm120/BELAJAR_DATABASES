package BELAJAR_DATABASES

import (
	"context"
	"fmt"
	"testing"
	"time"
	"database/sql"
	"strconv"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	var script[3]string

	script[0] = "INSERT INTO belajar_database(id, name, email, balance, rating, birth_date, married) VALUES ('ikbal', 'ikbal', 'ikbal@gmail.com', 1000000, 5.0, '2006-05-19',false)"
	script[1] = "INSERT INTO belajar_database(id, name, email, balance, rating, birth_date, married) VALUES ('kontol', 'kontol', 'kontol@gmail.com', 1000000, 5.0, '1987-09-28',true)"
	script[2] = "INSERT INTO belajar_database(id, name, email, balance, rating, birth_date, married) VALUES ('memek', 'memek', 'memek@gmail.com', 1000000, 5.0, '2006-01-04',false)"

	for i := 0; i < len(script); i++ {
		_, err := db.ExecContext(ctx, script[i])

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Success insert new belajar_database")
}


func TextExecSqlUpdate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "UPDATE belajar_database SET email = null, birth_date = null, rating = 8.3 WHERE id = 'memek' "
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("succes insert belajar database")
}


func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM belajar_database"
	rows, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
		}
	defer rows.close()

	for rows.Next() {
		var id, name string
	err = rows.Scan(&id, &name)
	if err != nil {
		panic(err)
	}
	fmt.Println("id:", id)
	fmt.Println("name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM belajar_database"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		// Posisi harus sesuai dengan mapping table
		if err != nil {
			panic(err)
		}

		fmt.Println("===============")
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if email.Valid {
			fmt.Println("Birth Date : ", birthDate)
		}
		fmt.Println("Created At : ", createdAt)
		fmt.Println("married : ", married)
	}
}

func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin" // Input user yang seharusnya
	username := "admin '; #" // Input user dengan SQL Injection
	password := "salah" // Password yg benar admin@123

	script := "SELECT username FROM user WHERE username = '"+username+
	"' AND password='"+password+"' LIMIT 1 "
	rows, err := db.QueryContext(ctx, script)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login : ", username)
	}else{
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin" // Input user yang seharusnya
	username := "admin" // Input user dengan SQL Injection
	password := "admin@123" // Password yg benar admin@123

	script := "SELECT username FROM user WHERE username = ? AND password=? LIMIT 1 "
	rows, err := db.QueryContext(ctx, script, username, password)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login : ", username)
	}else{
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	var script  string

	username := "ikbal"
	password := "ikbal@123"

	script = "INSERT INTO user(username, password) VALUES (?,?)"
	
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	var script  string

	email := "kontol@gmail.com"
	comment := "Test comment kontol"

	script = "INSERT INTO comments(email, comment) VALUES (?,?)"
	
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err :=  result.LastInsertId()
	if err != nil{
		panic(err)
	}

	fmt.Println("Success insert new comment with id :", insertId)
}

func TestPrepareStatement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx :=  context.Background()
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	statement, err :=  db.PrepareContext(ctx, script)
	if err != nil{
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "ikbal"+strconv.Itoa(i)+"@gmail.com"
		comment := "Komentar ke "+strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil{
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("Comment Id : ", id)
	}
}

func TestTransaction(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	// Do Transaction
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"

	for i := 9; i < 20; i++ {
		email := "kontol"+strconv.Itoa(i)+"@gmail.com"
		comment := "Komentar ke "+strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil{
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("Comment Id : ", id)
	}

	// End Transaction
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}

/**
 * EKSEKUSI PERINTAH SQL
 * Untuk menjalankan perintah SQL di golang bisa menggunakan function
 * (DB) ExectContext(context, sql, params)
 * params sifatnya tidak wajib
 * Ketika kita menggunakan perintah SQL, kita butuh menjalankan context, 
 * dengan context kita bisa mengirim sinyal cancel jika kita ingin mebatalkan
 * pengiriman perintah SQLnya
 * 
 * QUERY SQL
 * Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah 
 * Exec, namun jika kita membutuhkan hasil seperti SELECT SQL, kita bisa menggunakan
 * function berbeda yang berbeda 
 * Function untuk melakukan query ke databale, bisa menggunakan function 
 * (DB) QueryContext(context, sql, params) params tidak wajib
 * * ROWS
 * Hasil query function adalah sebuah data struct sql.Rows
 * Rows digunakan untuk melakukan iterasi terhadap hasil dari query 
 * Kita bisa menggunakan function (Rows)Next()(boolean) untuk melakukan iterasi terhadap 
 * data hasil query, jika return false, artinya sudah tidak ada data lagi dalam result
 * Untuk membaca tiap data kita bisa menggunakan (Rows)scan(columns...)
 * Dan jangan lupa, setelah menggunakan Rows, Jangan lupa untuk menutupnya menggunakan 
 * (rows)Close()
 * 
 * TIPE DATA COLUMN
 * * Mapping Tipe Data
 * VARCHAR, CHAR => string
 * INT, BIGINT => int32, int64
 * FLOAT, DOUBLE => float32, float64
 * BOOLEAN => bool
 * DATE, DATETIME, TIME, TIMESTAMP => time.Time
 * Saat select data disarankan menybutkan nama kolomnya agar tidak berubah-rubah saat 
 * ALTER TABLE dan tidak sarankan menggunakan SELECT * FROM table karena jika table dirubah
 * maka posisi di golang juga akan berubah
 * 
 * * Error Tipe Data Date
 * Message : "unsported Scan, storing driver.Value type[]uint8"
 * Secara default, Driver MySQL untuk golang akan melakukan query type data DATE, DATETIME
 * TIMESTAMP mejadi []byte / []uint8, dimana ini bisa dikonversi memjadi string lalu di 
 * parsing menjadi time.Time
 * Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta driver MySQL untuk golang
 * secara otomatis melakukan parsing dengan dengan menambahkan paramater parseTime=true
 * pada (DB)connection
 * 
 * * Nullable Type
 * Golang database tidak mengerti dengan type data NULL di database
 * Oleh karena itu, khusus untuk kolom yang bisa null di database, akan menjadi masalah jika
 * kita melakukan scan secara bulat-bulat menggunakan type data representasinya di golang
 * 
 * * Error Data Null
 * Message : converting NULL to string is unsported
 * Konversi secara otomatis tidak didukung oleh driver MySQL Golang
 * oleh karena itu, khususk kolom yang bisa null, kita perlu menggunakan type data yang ada
 * pada package sql
 * 
 * * Type Data Nullable
 * string => database/sql.NullString
 * bool => database/sql.NullBool
 * float64 => database/sql.NullFloat64
 * int32 => database/sql.NullInt32
 * int64 => database/sql.NullInt64
 * Time.Time => database/sql.NullTime
 * Data yang akan dikembalikan berupa struct 
 * 
 * SQL INJECTION
 * SQL Injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi 
 * dalam lapisan basis data sebuah aplikasi 
 * Biasanya, SQL Injection dilakukan dengan mengirim input dari user dengan perintah yang salah,
 * sehingga menyebabkan hasil SQL yang kita buat jadi tidak valid
 * SQL Injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman 
 * 
 * * Solusinya
 * Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat 
 * Jika kita membutuhkan paramter ketika membuat SQL, maka bisa menggunakan function Execute
 * atau query dengan parameter
 * 
 * SQL DENGAN PARAMETER
 * Sekarang kita sudah tahu bahayanya SQL Injection jika menggabungkan string ketika membuat query
 * Jika ada kebutuhan seperti itu, sebenarnya function exec dan query memiliki paramter tambahan
 * yang bisa digunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat
 * Untuk menandai sebuah SQL Membutuhkan paramter, Kita bisa menggunakan karakter ? (tanda tanya)
 * * Contoh SQL Injection
 * SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
 * INSERT INTO user(username, password) VALUES(?,?)
 * Dan lain-lain
 * 
 * AUTO INCREMENT
 * Kadang kita membuat sebuah table dengan id auto increment 
 * Dan kadang pula, kita ingin mengambil data yang sudah kita insert kedalam MySQL
 * Sebenarnya kita bisa melakukan query ulang dengan menggunakan SELECT LAST_INSERT_ID
 * namun tersebut cukup merepotkan karena harus 2x query
 * Untungnya di golang ada cara yang lebih mudah 
 * Kita bisa menggunakan function (result)LastInsertId() untuk mendapatkan Id terakhir 
 * yang dibuat secara auto increment
 * Result adalah object yang dikembalikan ketika kita menggunakan function Exec
 * 
 * PREPARE STATEMENT
 * Saat menggunakan Function Query atau Exec yang menggunakan parameter, sebenarnya 
 * implementasi dibawahnya menggunakan Prepare statement
 * Jadi tahapan pertama statementnya disiapkan terlebih dahalu, setelah itu baru 
 * diisi dengan parameter
 * Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya
 * berbeda parameternya saja, misalnya insert data langsung banyak
 * Pembuatan prepare statement bisa dilakukan dengan manual, tanpa harus menggunakan
 * Query atau Exec dengan parameter
 * 
 * * Prepare Statement
 * Saat membuat prepare statement, secara otomatis akan mengenali koneksi database
 * yang digunakan
 * Sehingga ketika kita mengekseskui Prepare statement berkali-kali, maka akan 
 * menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statementnya
 * hanya sekali diawal saja
 * Jika menggunakan Query atau Exec dengan parameter, kita tidak bisa menjamin bahwa 
 * koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepare statement 
 * akan selalu dibuat berkali-kali walaupun menggunakan SQL yang sama 
 * Untuk membuat Prepare Statement bisa bisa menggunakan (DB)Prepare(Contex, sql)
 * Prepare Statement direpresentasikan dalam struct database/sql.Stmt
 * Sama seperti resource sql lainnya, Stmt() harus di close jika sudah tidak digunakan
 * 
 * DATABASE TRANSACTION
 * Secara default, semua perintah SQL yang kita kirim menggunakan Golang, akan otomatis
 * di commit, atau istilahnya auto commit
 * Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara
 * otomatis di commit ke database
 * Untuk memulai transaksi kita bisa menggunakan function (DB)Begin(), dimana akan 
 * menghasilkan struct Tx yang merupakan representasi Transaction
 * Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana
 * hampir function di DB ada di Tx seperti Exec, Query Atau Prepare
 * Setelah selesai proses transaksi kita bisa menggunakan function (Tx)Commit()untuk 
 * melakukan commit atau Rollback() jika ingin membatalkan proses transaction
 */