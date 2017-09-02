package common

import "github.com/dtrinh100/Music-Playlist/src/api/db"

type Str2MapStr map[string](map[string]string)
type ErrMap map[string]string

// ErrorList holds a list of errors to return to the client in JSON-format.
type ErrorList struct {
	Errors ErrMap    `json:"errors"`
}

// ServerConfig holds the server's configuration settings.
type ServerConfig struct {
	Address  string
	LogLevel int
}

// Env holds the env (database, sessions, etc.) for usage by all handlers.
type Env struct {
	DB *db.DB
}
