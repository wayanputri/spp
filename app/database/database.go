package database

import (
	"database/sql"
	"fmt"
	"log"
	"project/app/config"

	_ "github.com/lib/pq"
)

func InitPostgres(cfg *config.AppConfig) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success connect database")
	}
	return db
}
