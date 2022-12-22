package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var dbconn *gorm.DB
func ConnectDB() *gorm.DB {

	errenv:=godotenv.Load()
	if errenv !=nil{
		log.Panic(errenv.Error())
	}

	DB_HOST:=os.Getenv("DB_HOST")
	DB_NAME:=os.Getenv("DB_NAME")
	DB_USER:=os.Getenv("DB_USER")
	DB_PASS:=os.Getenv("DB_PASS")
	DB_PORT:=os.Getenv("DB_PORT")
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":"+DB_PORT+")/" + DB_NAME + "?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		fmt.Println("Koneksi Gagal "+err.Error())
		
	}

	dbconn = database
	return dbconn
}
