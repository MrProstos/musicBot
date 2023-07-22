package Database

import (
	"MrProstos/download_utils/App/Utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// initConnection Initialize Database connection
func initConnection() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		Utils.GetEnv("DB_HOST"),
		Utils.GetEnv("DB_USER"),
		Utils.GetEnv("DB_PASSWORD"),
		Utils.GetEnv("DB_NAME"),
		Utils.GetEnv("DB_PORT"),
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = gormDB
}

// GetDB Get gorm.DB instance
func GetDB() *gorm.DB {
	if db == nil {
		initConnection()
	}

	return db
}

// RunMigration Run Database migration
func RunMigration(models ...interface{}) {
	err := GetDB().AutoMigrate(models)
	if err != nil {
		log.Fatalln(err)
	}
}
