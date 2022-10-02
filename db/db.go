package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
    Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
    db, err := gorm.Open(sqlite.Open("rest.db"), &gorm.Config{})
    
    if err != nil {
        log.Fatal("Failed to connec to the Database! \n", err.Error())
        os.Exit(2)
    }

    log.Println("Connected to DB successfully")
    db.Logger = logger.Default.LogMode(logger.Info)
    log.Println("Running Migrations")
    // Add migrations

    Database = DbInstance{Db: db}
}
