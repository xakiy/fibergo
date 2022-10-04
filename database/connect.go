package database

import (
    "fmt"
    "log"
    "strconv"

    "github.com/xakiy/fibergo/config"
    "github.com/xakiy/fibergo/internal/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err :=strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection  URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASS"), config.Config("DB_NAME"))
	// Connect to the DB and initiate the DB variables
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Connection to database failed!")
	}

	fmt.Println("Connection to database success!")

	// Do some DB migrations
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database migrated...")
}