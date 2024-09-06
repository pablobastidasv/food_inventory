package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func WithTransaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Info("Running with transaction")
		return next(c)
	}
}
