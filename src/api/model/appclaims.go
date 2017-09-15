package model

import "github.com/dgrijalva/jwt-go"

// AppClaims is used to store some JWT data
type AppClaims struct {
	jwt.StandardClaims
	RefreshAt int64  `json:"rat,omitempty"`
	UserEmail string `json:"email"`
}
