package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"test/config"
)

func NewSellerDatabase() *SellerDatabase {
	cfg := config.LoadENV(".env")

	log.Info().Msg("host: " + cfg.DBHost)
	log.Info().Msg("Postgres host: " + cfg.PostgresHost)

	//fmt.Println("name: ", cfg.DBName)
	//fmt.Println("port: ", cfg.DBPort)
	//fmt.Println("pass: ", cfg.UserDBPassword)
	//fmt.Println("driver: ", cfg.DriverDBName)
	//fmt.Println("username: ", cfg.UserDBName)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.DBPort, cfg.UserDBName, cfg.UserDBPassword, cfg.DBName)

	log.Info().Msgf("Connection string: %s", connStr)

	db, err := sql.Open(cfg.DriverDBName, connStr)
	if err != nil {
		log.Warn().Err(err).Msg("Unable to connect to database")
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to ping the database")
		return nil
	}
	log.Info().Msg("Successfully connected to the database.")

	return &SellerDatabase{Connection: db}
}

func (db *SellerDatabase) Close() {
	if db.Connection != nil {
		db.Connection.Close()
		log.Info().Msg("Database connection closed.")
	}
}
