package auth

import (
	"time"

	"github.com/MateusHBR/mallus_api/server"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	createAccount(CreateAccount) (string, error)
}

type AuthService struct {
	repository repository
}

func createToken(data CreateAccount) (string, error) {
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

	//TODO: get from os.Getenv("ACCESS_SECRET")
	tokenString, err := token.SignedString([]byte("my-secret-jwt-key"))

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

	token, err := createToken(res)

	if err != nil {
		return "", err
	}

	return token, nil
}
