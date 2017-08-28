package model

/*User represents the users using the site */
type User struct {
	Username       string `json:"username" bson:"username"`
	Password       string `json:"password,omitempty" bson:"-"`
	HashedPassword []byte `json:",omitempty" bson:"hashedpassword"`
	Email          string `json:"email" bson:"email"`
}
