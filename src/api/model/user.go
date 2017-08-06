package model

/*User represents the users using the site */
type User struct {
	Username string `json:"username" bson:"username"`
	Password []byte `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}
