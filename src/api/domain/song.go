package domain

import (
	"time"
)

type Song struct {
	ID          int           `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Artist      string        `json:"artist" bson:"artist"`
	Description string        `json:"description" bson:"description"`
	AudioPath   string        `json:"audiopath" bson:"audiopath"`
	ImgURL      string        `json:"imgurl" bson:"imgurl"`
	AltText     string        `json:"alttext" bson:"alttext"`
	Country     string        `json:"country" bson:"country"`
	State       string        `json:"state" bson:"state"`
	Duration    time.Duration `json:"duration" bson:"duration"`
	Likes       int           `json:"likes" bson:"likes"`
	Dislikes    int           `json:"dislikes" bson:"dislikes"`
}
