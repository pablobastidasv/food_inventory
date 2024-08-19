package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

func GetProductsNew() echo.HandlerFunc {
	return func(c echo.Context) error {
		return Render(c, http.StatusOK, pages.CreateProductPage())
	}
}

func PostProducts(pc inventorymanager.ProductCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
        slog.Debug("request to create product has arrived")
		product, err := pc.CreateProduct(
            c.Request().Context(),
			inventorymanager.CreateProductInput{
				Id:       c.FormValue("id"),
				Name:     c.FormValue("name"),
				Category: c.FormValue("category"),
			},
		)
		if err != nil {
			slog.Error("error creating product", "err", err)
			return err
		}

		return c.String(201, fmt.Sprintf("product created %v", product))
	}
}
