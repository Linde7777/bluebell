package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte("Red Read Redemption II")
var TokenExpireDuration = time.Hour * 2
var ErrInvalidToken = errors.New("invalid token")

type MyClaims struct {
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func GenToken(userID int64, username string) (string, error) {
	mc := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
		UserID:   userID,
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, err
	}
	return nil, ErrInvalidToken
}
