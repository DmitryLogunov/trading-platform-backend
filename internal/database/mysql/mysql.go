package mysql

import (
	"fmt"
	"github.com/DmitryLogunov/trading-platform/internal/database/mysql/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DBInstance a variable to store database connection
var DBInstance *gorm.DB

// Var for error handling
var err error

// ConnectDB connecting to the db
func ConnectDB() {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	fmt.Println(connectionString)

	// pass the db connection string
	connectionURI := connectionString
	// check for db connection
	DBInstance, err = gorm.Open("mysql", connectionURI)
	if err != nil {
		fmt.Println(err)
		// if the connection was unsuccessful
		panic("Database connection attempt was unsuccessful.....")
	} else {
		// if the connection was successful
		fmt.Println("Database Connected successfully.....")
	}
	// log all database operations performed by this connection
	DBInstance.LogMode(true)
}

func CreateDB() {
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// Create a database
	DBInstance.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", mysqlDatabase))
	// make the database available to this connection
	DBInstance.Exec(fmt.Sprintf("USE %s", mysqlDatabase))
}

func MigrateDB() {
	// migrate and sync the models to create a db table
	DBInstance.AutoMigrate(&models.Post{})
	fmt.Println("Database migration completed....")
}
