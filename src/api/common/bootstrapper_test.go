package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

/**
TestInitServer tests the initialization of ServerConfig (the
configuration settings for the server).
*/
func TestInitServer(t *testing.T) {
	asrt := assert.New(t)

	envVars := map[string]string{
		serverAddressKey: "0.0.0.0:8080",
		logLvlKey:        "4",
	}

	for envKey, envVal := range envVars {
		os.Setenv(envKey, envVal)
	}

	loglvl, logLvlErr := strconv.Atoi(envVars[logLvlKey])
	asrt.NoError(logLvlErr)

	expected := &ServerConfig{
		Address:  envVars[serverAddressKey],
		LogLevel: loglvl,
	}
	result := InitServer()

	asrt.Equal(expected, result)
}
