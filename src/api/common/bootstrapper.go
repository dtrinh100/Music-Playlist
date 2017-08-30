package common

import (
	"os"
	"strconv"
)

const (
	serveraddress_envkey = "MP_SRVRADDR_ENV"
	loglevel_envkey      = "MP_LOGLVL_ENV"
)

func InitServer() *ServerConfig {
	logLevelInt, lli_err := strconv.Atoi(os.Getenv(loglevel_envkey))
	if lli_err != nil {
		Fatal(lli_err, "Failed to convert "+os.Getenv(loglevel_envkey)+" to int.")
	}

	return &ServerConfig{
		Address:  os.Getenv(serveraddress_envkey),
		LogLevel: logLevelInt,
	}
}
