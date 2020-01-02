package config

import (
	"os"
)

func GetDatabaseUrl() string {
	return os.Getenv("UMPISA_DB_URL")
}
