package middleware

import (
	"context"
	"strings"

	errorhandler "bitbucket.org/cloud-platform/vnpt-sso-authentication/model/error_handler"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	"github.com/labstack/echo/v4"
)

// SSOValidateAuthenToken - SSOValidateAuthenToken
func SSOValidateAuthenToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		if c.Request().Header.Get("Authorization") == "" || len(strings.Split(c.Request().Header.Get("Authorization"), " ")) < 2 {
			return errorhandler.ErrorReturn(c, "1", "Authentication Token is Invalid", nil, nil)
		}
		accessToken := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]

		redisKey := "SSOToken$" + accessToken
		if !redis.Exists(ctx, redisKey) {
			return errorhandler.ErrorReturn(c, "1", "Authentication Token is Invalid", nil, nil)
		}

		ssoUserID, err := redis.Get(ctx, redisKey)
		if err != nil {
			return errorhandler.ErrorReturn(c, "10", "Failed: "+err.Error(), nil, nil)
		}

		collection := "SSO_USERS"
		if tempApp := mgodb.FindOneByField(initial.MgoDBName, collection, "user_id", ssoUserID); tempApp == nil {
			return errorhandler.ErrorReturn(c, "1", "Authentication Token is Invalid", nil, nil)
		}
		c.Request().Header.Set("user_id", ssoUserID)
		return next(c)
	}
}
