package domain

import (
	"time"
)

type Song struct {
	Name      string        `json:"name" bson:"name"`
	Duration  time.Duration `json:"duration" bson:"duration"`
	AudioPath string        `json:"audiopath" bson:"audiopath"`
	Likes     int           `json:"likes" bson:"likes"`
	Dislikes  int           `json:"dislikes" bson:"dislikes"`
}
