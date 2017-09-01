package common

import "github.com/dtrinh100/Music-Playlist/src/api/db"

type Str2mapstr map[string](map[string]string)
type ErrMap map[string]string

type ErrorList struct {
	Errors ErrMap    `json:"errors"`
}

type ServerConfig struct {
	Address  string
	LogLevel int
}



type Env struct {
	DB *db.DB
}
