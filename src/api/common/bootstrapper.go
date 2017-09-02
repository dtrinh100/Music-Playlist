package common

import (
	"os"
	"strconv"
)

const (
	serverAddressKey = "MP_SRVRADDR_ENV"
	logLvlKey        = "MP_LOGLVL_ENV"
)

/**
	InitServer helps initialize the server's configuration.
*/
func InitServer() *ServerConfig {
	logLevelInt, strConvErr := strconv.Atoi(os.Getenv(logLvlKey))
	if strConvErr != nil {
		Fatal(strConvErr, "Failed to convert \""+os.Getenv(logLvlKey)+"\" to int.")
	}

	return &ServerConfig{
		Address:  os.Getenv(serverAddressKey),
		LogLevel: logLevelInt,
	}
}
