package model

import "gopkg.in/mgo.v2/bson"

/*User represents the users using the site */
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Email    string        `json:"email" bson:"email"`
}
