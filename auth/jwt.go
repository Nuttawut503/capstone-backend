package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var (
	access_lifetime  = time.Minute * 15
	refresh_lifetime = time.Hour * 6
)

func createAccessToken(accessUUID, userID string, exp int64) (accessToken string, err error) {
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"authorized":  true,
		"access_uuid": accessUUID,
		"user_id":     userID,
		"exp":         exp,
	}).SignedString([]byte(viper.GetString("ACCESS_SECRET")))
	return
}

func createRefreshToken(refreshUUID, userID string, exp int64) (refreshToken string, err error) {
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"refresh_uuid": refreshUUID,
		"user_id":      userID,
		"exp":          exp,
	}).SignedString([]byte(viper.GetString("REFRESH_SECRET")))
	return
}

func extractToken(header http.Header) (string, error) {
	authToken := header.Get("Authorization")
	splitted := strings.Split(authToken, " ")
	if len(splitted) != 2 {
		return "", errors.New("Wrong format")
	}
	return splitted[1], nil
}

func verifyToken(token, secret string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	return t, err
}

func isValid(t *jwt.Token) bool {
	_, ok := t.Claims.(jwt.MapClaims)
	return ok && t.Valid
}
