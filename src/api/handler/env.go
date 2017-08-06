package handler

import "github.com/dtrinh100/Music-Playlist/src/api/db"

// Env holds the env (database, sessions, etc.)
// for usage by all handlers.
type Env struct {
	DB *db.DB
}
