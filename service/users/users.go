package users

import (
	"context"
	"fmt"
	"strings"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/handler/usershandler"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/model"
	errorhandler "bitbucket.org/cloud-platform/vnpt-sso-authentication/model/error_handler"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/utility"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	// ctx := context.Background()
	userProfile := new(model.UserProfile)
	err := c.Bind(&userProfile)
	if err != nil {
		return errorhandler.ErrorReturn(c, "6", "Body Invalid", nil, nil)
	}

	collection := "SSO_USERS"
	if tempUser := mgodb.FindOneByField(initial.MgoDBName, collection, "username", userProfile.Username); tempUser != nil {
		return errorhandler.ErrorReturn(c, "2", "Registration Failed. User is Exists", nil, nil)
	}

	userID := utility.GenerateUserID("SSOUSERS", collection)
	if userID == "" {
		return errorhandler.ErrorReturn(c, "10", "Registration Failed1", nil, nil)
	}

	if n, err := mgodb.CountDocumentByField(initial.MgoDBName, collection, "phone_number", userProfile.PhoneNumber); err != nil || n > 0 {
		return errorhandler.ErrorReturn(c, "2", "Register Failed. You need use another Phone Number", nil, nil)
	}
	if n, err := mgodb.CountDocumentByField(initial.MgoDBName, collection, "email", userProfile.Email); err != nil || n > 0 {
		return errorhandler.ErrorReturn(c, "2", "Register Failed. You need use another Email", nil, nil)
	}

	if userProfile.Password, err = utility.Hash(userProfile.Password); err != nil {
		return errorhandler.ErrorReturn(c, "10", "Registration Failed2", nil, nil)
	}

	if err := mgodb.SaveMongo(initial.MgoDBName, collection, userID, userProfile); err != nil {
		fmt.Println(err)
		return errorhandler.ErrorReturn(c, "10", "Registration Failed3", nil, nil)
	}

	return errorhandler.ErrorReturn(c, "0", "OK", nil, nil)
}

func Login(c echo.Context) error {
	ctx := context.Background()

	dataLogin := new(model.DataLogin)
	err := c.Bind(dataLogin)
	if err != nil {
		return errorhandler.ErrorReturn(c, "8", "Username or Password Invalid", nil, nil)
	}

	if dataLogin.Username == "" || dataLogin.Password == "" {
		return errorhandler.ErrorReturn(c, "8", "Username or Password Invalid", nil, nil)
	}

	data, code := usershandler.LoginHandler(ctx, dataLogin.Username, dataLogin.Password)
	switch code {
	case 0:
		return errorhandler.ErrorReturn(c, "0", "OK", data, nil)
	case -1:
		return errorhandler.ErrorReturn(c, "8", "Username or Password is Invalid", nil, nil)
	case -2:
		return errorhandler.ErrorReturn(c, "10", "Generate Token Failed", nil, nil)
	case -3:
		return errorhandler.ErrorReturn(c, "4", "User is not Exist", nil, nil)
	default:
		return errorhandler.ErrorReturn(c, "10", "Login Failed", nil, nil)
	}
}

func ValidateSSOAccessToken(c echo.Context) error {
	ctx := context.Background()
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || len(strings.Split(authorizationHeader, " ")) < 2 {
		return errorhandler.ErrorReturn(c, "1", "Authentication Token Invalid", nil, nil)
	}
	accessToken := strings.Split(authorizationHeader, " ")[1]
	code := usershandler.ValidateSSOAccessToken(ctx, accessToken)
	if code != 0 {
		return errorhandler.ErrorReturn(c, "1", "Authentication Token Invalid", nil, nil)
	}
	return errorhandler.ErrorReturn(c, "0", "OK", nil, nil)
}
