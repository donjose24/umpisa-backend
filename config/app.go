package config

import (
	"os"
)

func GetApplicationKey() string {
	return os.Getenv("UMPISA_APPLICATION_KEY")
}
