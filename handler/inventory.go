package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/views/components"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

func GetProductsNew(pl inventorymanager.ProductsLister) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []components.ProductOverview

		list, err := pl.ListProduct(c.Request().Context())
		if err != nil {
			return err
		}

		slog.Info("products to list", "products", list)
		products = []components.ProductOverview{}
		for _, p := range list {
			products = append(products, components.ProductOverview{
				Id:       p.Id,
				Name:     p.Name,
				Category: p.Category.Name,
			})
		}

		return Render(c, http.StatusOK, pages.CreateProductPage(products))
	}
}

func PostProducts(pc inventorymanager.ProductCreator) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Debug("request to create product has arrived")
		_, err := pc.CreateProduct(
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

		return Render(c, http.StatusCreated, components.ProductForm())
	}
}
