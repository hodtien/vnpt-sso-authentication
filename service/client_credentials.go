package service

import (
	"context"
	"fmt"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/oauth2"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/handler"

	"github.com/labstack/echo/v4"
)

func CreateCredential(c echo.Context) error {
	ctx := context.Background()
	userID := c.Request().Header.Get("userid")

	ret, err := handler.SaveCredentialsToDB(ctx, userID)
	if err != nil {
		return c.JSON(400, map[string]interface{}{"code": "0", "message": "Create Credential Failed: " + err.Error()})
	}

	return c.JSON(200, map[string]interface{}{"code": "0", "message": "OK", "data": ret})
}

func Token(c echo.Context) error {
	ctx := context.Background()
	a, _ := oauth2.Srv.Manager.GetClient(ctx, "086f7b6f-31f7-43d6-8360-9a06fd4ccbe6")
	fmt.Println(a)
	return oauth2.Srv.HandleTokenRequest(c.Response().Writer, c.Request())
}

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := oauth2.Srv.ValidationBearerToken(c.Request())
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(400, map[string]interface{}{"code": "1", "message": err.Error()})
		}

		fmt.Println("Validate token OK", err)
		return next(c)
	}
}

func Test(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "OK"})
}
