package server

import (
	"github.com/labstack/echo/v4"
)

func WithTransaction(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        c.Logger().Info("Running with transaction")
        return next(c)
    }
}
