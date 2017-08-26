package common

import (
	"os"
	"strconv"
)

const (
	server_envkey    = "MP_SERVER_ENV"
	database_envkey  = "MP_DB_ENV"
	log_level_envkey = "MP_LOGLVL_ENV"
)

type configuration struct {
	Server, Database string
	LogLevel         int
}

var AppConfig configuration

/**
Note: the code in this block could have been placed in init() but
	it was placed here instead for self-documenting-code purposes.
**/
func InitServer() {
	loadAppConfig()
}

func loadAppConfig() {
	if logLevelInt, lliErr := strconv.Atoi(os.Getenv(log_level_envkey)); lliErr == nil {
		AppConfig = configuration{
			Server:   os.Getenv(server_envkey),
			Database: os.Getenv(database_envkey),
			LogLevel: logLevelInt,
		}
	} else {
		Fatal(lliErr, "Failed to convert "+os.Getenv(log_level_envkey)+" to int.")
	}
}
