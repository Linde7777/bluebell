package jwt

import (
	"bluebell/settings"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secret = []byte("Red Read Redemption II")
var ErrInvalidToken = errors.New("invalid token")

type MyClaims struct {
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func GenToken(userID int64, username string) (accessToken, refreshToken string, err error) {
	// do not move the following two lines of code outside of this
	// function, since before this function was called,
	// the setting.Conf is empty.
	var AccTokenExpDur = time.Duration(
		settings.Conf.AccessTokenExpireDuration) * time.Second
	var RefTokenExpDur = time.Duration(
		settings.Conf.RefreshTokenExpireDuration) * time.Second
	mc := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccTokenExpDur).Unix(),
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
		ExpiresAt: time.Now().Add(RefTokenExpDur).Unix(),
		Issuer:    "bluebell",
	})
	refreshToken, err = token.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return
}

// ParseToken can parse both accessToken and refreshToken,
// because refreshToken is represented in jwt.StandardClaims,
// and AccessToken, which represent in MyClaims, contain
// jwt.StandardClaims too
func ParseToken(tokenStr string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, KeyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, ErrInvalidToken
}

func IsTimeExpireErr(err error) bool {
	v, _ := err.(*jwt.ValidationError)
	return v.Errors == jwt.ValidationErrorExpired
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
