package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

//
//import (
//	"github.com/caarlos0/env/v9"
//	"github.com/joho/godotenv"
//	"github.com/rs/zerolog/log"
//)
//
//type Config struct {
//	UserDBPassword string `env:"DB_PASSWORD"`
//	UserDBName     string `env:"DB_USER"`
//	DBName         string `env:"DB_NAME"`
//	DriverDBName   string `env:"DB_DRIVER"`
//	DBHost         string `env:"DB_HOST"`
//	DBPort         string `env:"DB_PORT"`
//	URL            string `env:"URL"`
//	ApiKey         string `env:"API_KEY"`
//	Email          string `env:"EMAIL"`
//	EmailPassword  string `env:"EMAIL_PASSWORD"`
//}
//
//func LoadENV(filename string) *Config {
//	err := godotenv.Load(filename)
//	if err != nil {
//		log.Panic().Err(err).Msg(" does not load .env")
//	}
//	log.Info().Msg("successfully load .env")
//	cfg := Config{}
//	return &cfg
//
//}
//
//func (cfg *Config) ParseENV() {
//
//	err := env.Parse(cfg)
//	if err != nil {
//		log.Panic().Err(err).Msg(" unable to parse environment variables")
//	}
//	log.Info().Msg("successfully parsed .env")
//}

type Config struct {
	UserDBPassword string `env:"POSTGRES_PASSWORD"`
	UserDBName     string `env:"POSTGRES_USER"`
	DBName         string `env:"POSTGRES_DB"`
	DriverDBName   string `env:"POSTGRES_DRIVER"`
	DBHost         string `env:"DB_HOST"`
	PostgresHost   string `env:"POSTGRES_HOST"`
	DBPort         string `env:"POSTGRES_PORT"`
}

var cfg *Config

func LoadENV(filename string) *Config {
	if cfg != nil {
		return cfg
	}

	err := godotenv.Load(filename)
	if err != nil {
		log.Panic().Err(err).Msg("Error loading .env file")
	}
	log.Info().Msg("Successfully loaded .env")

	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Panic().Err(err).Msg("Error parsing environment variables")
	}
	log.Info().Msg("Successfully parsed .env")

	return cfg
}
