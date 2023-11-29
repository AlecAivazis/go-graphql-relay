package config

// this file has such a weird name to ensure it loads before any other files
// so that .env can set environment variables before they get loaded into
// config values

import (
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
