package usershandler

import (
	"context"
	"fmt"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/utility"
)

// LoginHandler - LoginHandler
func LoginHandler(ctx context.Context, username, password string) (map[string]string, int) {
	collection := "SSO_USERS"
	tempUser := mgodb.FindOneByField(initial.MgoDBName, collection, "username", username)
	if tempUser == nil {
		return nil, -3
	}
	// Check in your db if the user exists or not
	usernameDB := tempUser["username"].(string)
	hashedPassword := tempUser["password"].(string)
	uID := tempUser["user_id"].(string)
	if usernameDB == username {
		if err := utility.CheckPasswordHash(hashedPassword, password); err != nil {
			return nil, -1
		}

		authenToken, refreshToken, err := GenerateTokenPair(ctx, uID)
		if err != nil {
			fmt.Println(err)
			return nil, -2
		}
		return map[string]string{
			"token":         authenToken,
			"refresh_token": refreshToken,
		}, 0
	}
	return nil, -1
}

func ValidateSSOAccessToken(ctx context.Context, accessToken string) int {
	redisKey := "SSOToken$" + accessToken
	if redis.Exists(ctx, redisKey) {
		return 0
	}
	return -1
}