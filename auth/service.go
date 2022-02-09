package auth

import (
	"fmt"
	"time"

	"github.com/MateusHBR/mallus_api/server"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	createAccount(CreateAccount) (string, error)
}

type AuthService struct {
	jwtKey     string
	repository repository
}

func (as AuthService) validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}

		return []byte(as.jwtKey), nil
	})
}

func (as AuthService) createToken(data CreateAccount) (string, error) {
	var expires = server.TimeNow().Add(time.Minute * 15)

	claims := &JwtClaims{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(as.jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (as AuthService) createAccount(user CreateAccount) (string, error) {
	res, err := as.repository.createAccount(user)

	if err != nil {
		return "", err
	}

	token, err := as.createToken(res)

	if err != nil {
		return "", err
	}

	return token, nil
}
