package database

import (
	"fmt"
	"log"
	"time"
	"webapi-with-go/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // Definido minusculo pois não vai ser exportavel

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB { // Método para  instaciar o banco sem precisar ficar abrindo varias conexões
	return db
}
