package tokenutil

import (
	"fmt"
	"strconv"
	"time"

	USER_MODEL "onlyfounds/module/user/model"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *USER_MODEL.User, secret string, expiry string) (accessToken string, err error) {
	exptime, err := strconv.Atoi(expiry)
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(time.Minute * time.Duration(exptime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserName,
		"exp": exp,
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *USER_MODEL.User, secret string, expiry string) (refreshToken string, err error) {
	exptime, err := strconv.Atoi(expiry)
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(time.Minute * time.Duration(exptime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserName,
		"exp": exp,
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractUserNameFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["sub"].(string), nil
}
