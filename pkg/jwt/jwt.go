package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte("Red Read Redemption II")
var AccTokenExpireDuration = time.Minute * 10
var RefTokenExpireDuration = time.Hour * 24 * 30
var ErrInvalidToken = errors.New("invalid token")

type MyClaims struct {
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func GenToken(userID int64, username string) (accessToken, refreshToken string, err error) {
	mc := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccTokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
		UserID:   userID,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	accessToken, err = token.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefTokenExpireDuration).Unix(),
		Issuer:    "bluebell",
	})
	refreshToken, err = token.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return
}

// ParseToken can parse both AccessToken and RefreshToken,
// because RefreshToken is represented in jwt.StandardClaims,
// and AccessToken, which represent in MyClaims, contain
// jwt.StandardClaims too
func ParseToken(tokenStr string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, KeyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, err
	}
	return nil, ErrInvalidToken
}

func RefreshToken(accToken, refToken string) (newAccToken, newRefToken string,
	err error) {
	if _, err = jwt.Parse(refToken, KeyFunc); err != nil {
		return
	}

	var mc MyClaims
	_, err = jwt.ParseWithClaims(accToken, &mc, KeyFunc)
	v, _ := err.(*jwt.ValidationError)

	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(mc.UserID, mc.Username)
	}

	return
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return secret, nil
}
