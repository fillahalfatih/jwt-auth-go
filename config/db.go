// config/db.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// HAPUS "var DB *gorm.DB". Kita tidak membutuhkannya lagi.

// UBAH NAMA FUNGSI DAN TAMBAHKAN RETURN TYPES (*gorm.DB, error)
func ConnectDB() (*gorm.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// Jangan gunakan log.Fatal di sini, cukup kembalikan errornya
		// Biarkan pemanggil (main.go) yang memutuskan untuk mem-fatal-kan aplikasi
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)

	// GUNAKAN VARIABEL LOKAL `db`, bukan global `DB`
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // Kembalikan error jika koneksi gagal
	}

	log.Println("Connected to MySQL database!")

	// KEMBALIKAN objek `db` dan `nil` untuk error jika sukses
	return db, nil
}