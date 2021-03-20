package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	_ "bitbucket.org/cloud-platform/vnpt-sso-authentication/oauth2"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/service"
	mdw "bitbucket.org/cloud-platform/vnpt-sso-authentication/middleware"
	users "bitbucket.org/cloud-platform/vnpt-sso-authentication/service/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	go terminateProcess()

	e := echo.New()
	s := &http.Server{
		Addr:         ":15001",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	e.Use(middleware.CORS())

	e.POST("/api/vnpt-sso-authentication/v1/credentials", service.CreateCredential)
	e.GET("/api/vnpt-sso-authentication/v1/get_token", service.Token, mdw.SSOValidateAuthenToken)
	e.GET("/test", service.Test)

	e.POST("/api/vnpt-sso-authentication/v1/user/register", users.Register)
	e.POST("/api/vnpt-sso-authentication/v1/user/login", users.Login)
	e.GET("/api/vnpt-sso-authentication/v1/user/validate_access_token", users.ValidateSSOAccessToken)

	e.Logger.Fatal(e.StartServer(s))
}

func terminateProcess() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal. Close All Redis Connection.")
		os.Exit(0)
	}
}
