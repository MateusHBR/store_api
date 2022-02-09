package auth

import "github.com/dgrijalva/jwt-go"

type CreateAccount struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type JwtClaims struct {
	ID        string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	jwt.StandardClaims
}
