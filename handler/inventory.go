package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/views/components"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

func GetInventory() echo.HandlerFunc{
    return func(c echo.Context) error {
        return Render(c, 200, pages.InventoryPage())
    }

}

func GetInventoryItem() echo.HandlerFunc{
    return func(c echo.Context) error {
        return Render(c, 200, components.InventoryValue())
    }
}
    
func GetInventoryForm() echo.HandlerFunc{
    return func(c echo.Context) error {
        return Render(c, 200, components.InventoryItemForm())
    }
}
