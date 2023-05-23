package database

import (
	"fmt"
	dbconfig "go-gin-api/config/dbConfig"
	"go-gin-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(){
  var err error

  if errEnv := godotenv.Load(); errEnv != nil {
    log.Fatal("Tidak bisa mengakses file .env")
  }  

  dsn := os.Getenv("DSN")
  dbconfig.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Gagal koneksi ke database")
  } else {
    fmt.Println("Berhasil terkoneksi ke database")
  }
  
}

func RunMigration(){
  if err := dbconfig.DB.AutoMigrate(&models.User{}); err != nil {
    log.Fatal("Gagal migration database")
  } else {
    fmt.Println("Berhasil migration database")
  }
}
