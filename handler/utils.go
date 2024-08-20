package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/views/components"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func renderMessages(ctx echo.Context, messages []components.AlertMessage) error {
	t := components.Messages(messages)

    buf := templ.GetBuffer()
    defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(http.StatusOK, buf.String())

}

func RenderMessage(ctx echo.Context, level string, message string) error {
	msgs := []components.AlertMessage{
		{
			Level:   level,
			Message: message,
		},
	}
    return renderMessages(ctx, msgs)
}
