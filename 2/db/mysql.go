package database

import (
	// _ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var getDb *gorm.DB
var dbStock string

func init() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open("root:travelio123@tcp(localhost:3306)/stock"), &gorm.Config{})
	// db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)")
	if err != nil {
		panic("Cannot connect to database")
	}

	getDb = db
	dbStock = "stock"

}

func GetDB() gorm.DB {
	return *getDb
}

// var dbMaster *sqlx.DB
// var dbStock string

// func init() {

// 	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)")
// 	if err != nil {
// 		panic("Cannot connect to database")
// 	}

// 	dbMaster = db.Unsafe()
// 	dbStock = "stock"

// }

// func DBStock() string {
// 	return dbStock
// }
