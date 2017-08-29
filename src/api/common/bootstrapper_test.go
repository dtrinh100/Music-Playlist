package common

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
)

/**
	TestLoadAppConfig tests the initialization of AppConfig (the
	configuration settings for the server).
*/
func TestLoadAppConfig(t *testing.T) {
	assert := assert.New(t)

	env_vars := map[string]string{
		serveraddress_envkey: "0.0.0.0:8080",
		dbname_envkey:        "databasename",
		loglevel_envkey:      "4",
	}

	for envkey, envval := range env_vars {
		os.Setenv(envkey, envval)
	}

	loadAppConfig()

	loglvl, loglvl_err := strconv.Atoi(env_vars[loglevel_envkey])
	assert.NoError(loglvl_err)

	expected := configuration{
		Server: server{
			Address:  env_vars[serveraddress_envkey],
			LogLevel: loglvl,
		},
		DB: database{
			Name: env_vars[dbname_envkey],
			UserTable: table{
				Name: usertable_name,
			},
		},
	}
	result := AppConfig

	assert.Equal(result, expected)
}
