package token

import "github.com/dgrijalva/jwt-go"

//Jwt json web token
type Jwt struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
