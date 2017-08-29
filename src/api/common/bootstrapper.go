package common

import (
	"os"
	"strconv"
)

const (
	serveraddress_envkey = "MP_SRVRADDR_ENV"
	dbname_envkey        = "MP_DBNAME_ENV"
	loglevel_envkey      = "MP_LOGLVL_ENV"
)

const (
	usertable_name = "users"
)

var AppConfig configuration

/**
Note: the code in this block could have been placed in init() but
	it was placed here instead for self-documenting-code purposes.
**/
func InitServer() {
	loadAppConfig()
}

func loadAppConfig() {
	if logLevelInt, lliErr := strconv.Atoi(os.Getenv(loglevel_envkey)); lliErr == nil {
		srvr := server{
			Address:  os.Getenv(serveraddress_envkey),
			LogLevel: logLevelInt,
		}

		db := database{
			Name: os.Getenv(dbname_envkey),
			UserTable: table{
				Name: usertable_name,
			},
		}

		AppConfig = configuration{
			Server: srvr,
			DB:     db,
		}
	} else {
		Fatal(lliErr, "Failed to convert "+os.Getenv(loglevel_envkey)+" to int.")
	}
}
