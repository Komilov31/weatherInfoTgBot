package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Komilov31/weatherInfoBot/config"
	_ "github.com/lib/pq"
)

func InitStorage() *sql.DB {

	DBConfig := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	)

	db, err := NewSqlStorage(DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
	return db
}

func NewSqlStorage(cfg string) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
