package common

import (
	"os"
	"strconv"
)

const (
	serveraddress_envkey = "MP_SRVRADDR_ENV"
	loglevel_envkey      = "MP_LOGLVL_ENV"
)

/**
	InitServer helps initialize the server's configuration.
*/
func InitServer() *ServerConfig {
	logLevelInt, lli_err := strconv.Atoi(os.Getenv(loglevel_envkey))
	// TODO: Handle this more appropriately.
	if lli_err != nil {
		Fatal(lli_err, "Failed to convert "+os.Getenv(loglevel_envkey)+" to int.")
	}

	return &ServerConfig{
		Address:  os.Getenv(serveraddress_envkey),
		LogLevel: logLevelInt,
	}
}
