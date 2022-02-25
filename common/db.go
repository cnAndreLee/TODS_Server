package common

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// func InitDB() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmoe=disable", "localhost", 54321, "postgres", "postgres", "users")
// 	db, err := sql.Open("postgres", psqlInfo)

// 	if err != nil {
// 		return
// 	}

// 	DB = db
// }

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=tods port=54321 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
