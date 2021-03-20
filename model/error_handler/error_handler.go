package errorhandler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// ErrorReturn - ErrorReturn
func ErrorReturn(c echo.Context, code string, message interface{}, data interface{}, err error) error {
	if err != nil {
		fmt.Println(err)
	}
	if code != "0" {
		return c.JSON(400, map[string]interface{}{"code": code, "message": message, "data": data})
	}
	return c.JSON(200, map[string]interface{}{"code": code, "message": message, "data": data})
}

// ErrorReturnWithTotal - ErrorReturnWithTotal
func ErrorReturnWithTotal(c echo.Context, code, message string, data interface{}, total interface{}, err error) error {
	if err != nil {
		fmt.Println(err)
	}
	if code != "0" {
		return c.JSON(400, map[string]interface{}{"code": code, "message": message, "data": data, "total": total})
	}
	return c.JSON(200, map[string]interface{}{"code": code, "message": message, "data": data, "total": total})
}