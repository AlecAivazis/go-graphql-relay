package config

import (
	"os"
	"strconv"
)

var DbHost = "localhost"
var DbPort = 5432
var DbName = "development"
var DbUser = "ent"
var DbPassword = "password"
var DbEnableSSL = false

func init() {
	if e := os.Getenv("DB_HOST"); e != "" {
		DbHost = e
	}

	if e := os.Getenv("DB_PORT"); e != "" {
		val, err := strconv.ParseInt(e, 10, 64)
		if err == nil {
			DbPort = int(val)
		}
	}

	if e := os.Getenv("DB_NAME"); e != "" {
		DbName = e
	}

	if e := os.Getenv("DB_USER"); e != "" {
		DbUser = e
	}

	if e := os.Getenv("DB_PASSWORD"); e != "" {
		DbPassword = e
	}

	if n := os.Getenv("DB_ENABLE_SSL"); n == "true" {
		DbEnableSSL = true
	}
}
