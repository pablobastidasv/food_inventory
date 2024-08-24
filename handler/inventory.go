package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/views/components"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

func GetMainIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		return Render(c, 200, pages.InventoryPage())
	}

}

func GetInventoryItems(lister inventorymanager.InventoryItemsLister) echo.HandlerFunc {
	return func(c echo.Context) error {
        inventoryitems, err := lister.ListInventoryItems(c.Request().Context())
        if err != nil {
            return err
        }

        items := []components.InventoryInfo{}
        for _, i := range inventoryitems{
            items = append(items, components.InventoryInfo{
            	Id:          i.Id,
            	ProductName: i.Product.Name,
            	Ammount:     strconv.Itoa(i.Ammount),
            })
        }

		return Render(c, 200, components.InventoryItems(items))
	}
}

func GetInventoryForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return Render(c, 200, components.InventoryItemForm())
	}
}

func PutInventory() echo.HandlerFunc {
	return func(c echo.Context) error {
        id := c.Param("id")
		return Render(c, 200, components.InventoryValue(id, "13"))

	}
}
