package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv berfungsi untuk membaca .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetEnv berfungsi untuk mendapatkan nilai dari .env
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// MustGetEnv mengambil environment variable yang WAJIB diisi.
// Jika tidak diset, aplikasi langsung dihentikan (fatal) alih-alih
// diam-diam memakai nilai default yang tidak aman.
func MustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("FATAL: required environment variable %s is not set. Application cannot start without it.", key)
	}
	return value
}
