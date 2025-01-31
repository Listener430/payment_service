package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				c.String(http.StatusInternalServerError, "internal server error")
			}
		}()
		return next(c)
	}
}
