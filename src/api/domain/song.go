package domain

import (
	"time"
)

type Song struct {
	ID        int           `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	AudioPath string        `json:"audiopath" bson:"audiopath"`
	Country   string        `json:"country" bson:"country"`
	State     string        `json:"state" bson:"state"`
	Duration  time.Duration `json:"duration" bson:"duration"`
	Likes     int           `json:"likes" bson:"likes"`
	Dislikes  int           `json:"dislikes" bson:"dislikes"`
}
