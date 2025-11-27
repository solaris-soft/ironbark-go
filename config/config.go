package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
}

func LoadConfig() *Config {
	godotenv.Load()
	dbURL := flag.String("db", os.Getenv("DB_URL"), "Database URL")
	flag.Parse()
	return &Config{
		DBURL: *dbURL,
	}
}
