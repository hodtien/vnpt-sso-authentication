package usershandler

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateTokenPair - GenerateTokenPair
func GenerateTokenPair(ctx context.Context, ssoUserID string) (string, string, error) {
	// Initial
	jwtSecretKey := "Rn6q27RXd3zZBMhAdk83n"

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = ssoUserID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	authToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshTokenJWT.Claims.(jwt.MapClaims)
	rtClaims["sub"] = ssoUserID
	rtClaims["exp"] = time.Now().Add(time.Hour * 8760).Unix() //refresh_token 1 year expire

	refreshToken, err := refreshTokenJWT.SignedString([]byte(jwtSecretKey))
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	//redisKey := "VDP_SSO$" + ssoUserID
	redisKey := "SSOToken$" + authToken
	redisKeyRefToken := "SSOToken$RefreshToken$" + "$" + refreshToken
	if _, err := redis.SetExpire(ctx, redisKey, ssoUserID, 2*time.Hour); err != nil {
		redis.Delete(ctx, redisKey)
		return "", "", nil
	}
	if _, err := redis.SetExpire(ctx, redisKeyRefToken, ssoUserID, 8760*time.Hour); err != nil {
		redis.Delete(ctx, redisKey)
		return "", "", nil
	}

	return authToken, refreshToken, nil
}
