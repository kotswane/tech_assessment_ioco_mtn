package db

import (
	 "fmt"
	 "os"
	 "gorm.io/driver/mysql"
	 "gorm.io/gorm"
)

 var (

	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
 )

func Init() (*gorm.DB) {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt:true})
	return db
}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}
