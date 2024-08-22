package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/views/components"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

type PostProductDeps interface {
	inventorymanager.CategoryLister
	inventorymanager.ProductCreator
}

func GetProducts(pl inventorymanager.ProductsLister) echo.HandlerFunc {
	return func(c echo.Context) error {
		var products []components.ProductOverview

		list, err := pl.ListProduct(c.Request().Context())
		if err != nil {
			return err
		}

		products = []components.ProductOverview{}
		for _, p := range list {
			products = append(products, components.ProductOverview{
				Id:       p.Id,
				Name:     p.Name,
				Category: p.Category.Name,
			})
		}

		return Render(c, http.StatusOK, pages.ProductsPage(products))
	}
}

func PostProducts(d PostProductDeps) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Debug("request to create product has arrived")
		p, err := d.CreateProduct(
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

		po := components.ProductOverview{
			Id:       p.Id,
			Name:     p.Name,
			Category: p.Category.Name,
		}

		RenderMessage(c, "INFO", "Producto creado satisfactoriamente")
		renderProductForm(c, d)

		return Render(c, http.StatusCreated, components.ProductRowOob(po))
	}
}

func GetProductsForm(cl inventorymanager.CategoryLister) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderProductForm(c, cl)
	}
}

func renderProductForm(ctx echo.Context, cl inventorymanager.CategoryLister) error {
	categoryList, err := cl.ListCategories(ctx.Request().Context())
	if err != nil {
		return err
	}

	categories := []components.Category{}
	for _, c := range categoryList {
		category := components.Category{
			Id:   c.Code,
			Name: c.Name,
		}
		categories = append(categories, category)
	}

	return Render(ctx, http.StatusOK, components.ProductForm(categories))

}
