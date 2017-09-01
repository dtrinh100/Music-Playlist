package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
)

/**
	TestInitServer tests the initialization of ServerConfig (the
	configuration settings for the server).
*/
func TestInitServer(t *testing.T) {
	assert := assert.New(t)

	env_vars := map[string]string{
		serveraddress_envkey: "0.0.0.0:8080",
		loglevel_envkey:      "4",
	}

	for envkey, envval := range env_vars {
		os.Setenv(envkey, envval)
	}

		loglvl, loglvl_err := strconv.Atoi(env_vars[loglevel_envkey])
		assert.NoError(loglvl_err)

	expected := &ServerConfig{
		Address: env_vars[serveraddress_envkey],
		LogLevel: loglvl,
	}
	result := InitServer()

	assert.Equal(expected, result)
}
