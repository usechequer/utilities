package utilities

import (
	"fmt"
	"log"
	"os"
	"time"

	githubmysql "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

// GetDtabaseObject returns a struct pointer through which a MySQL
// database can be accessed
func GetDatabaseObject() *gorm.DB {
	if database != nil {
		return database
	}

	location, _ := time.LoadLocation("Europe/London")

	db, err := gorm.Open(mysql.New(mysql.Config{DSNConfig: &githubmysql.Config{User: os.Getenv("DATABASE_USERNAME"), Passwd: os.Getenv("DATABASE_PASSWORD"), DBName: os.Getenv("DATABASE_NAME"), Net: "tcp", Addr: fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT")), ParseTime: true, Loc: location}}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("There was a problem connecting to the database")
	}

	database = db

	return database
}
