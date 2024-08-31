package handler

import (
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/views/components"
	"github.com/pablobastidasv/fridge_inventory/views/pages"
)

type (
	putInventoryDeps interface {
		inventorymanager.InventoryItemUpdater
		inventorymanager.InventoryItemFinder
	}
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

		tmp := make(map[string]*components.InventoryCategory)
		categories := []components.InventoryCategory{}
		for _, i := range inventoryitems {
			cat, exists := tmp[i.Product.Category.Code]
			if !exists {
				cat = &components.InventoryCategory{
					CategoryName:  i.Product.Category.Name,
					Items:         []components.InventoryItemInfo{},
				}

				tmp[i.Product.Category.Code] = cat
			}

			itemInfo := components.InventoryItemInfo{
				Id:          i.Id,
				ProductName: i.Product.Name,
				Amount:      strconv.Itoa(i.Amount),
			}
			cat.Items = append(cat.Items, itemInfo)
		}

		for _, v := range tmp {
			categories = append(categories, *v)
		}

        sort.SliceStable(categories, func(i, j int) bool {
            return categories[i].CategoryName < categories[j].CategoryName
        })

		return Render(c, 200, components.InventoryItemsByCategory(categories))
	}
}

func GetInventoryForm(f inventorymanager.InventoryManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		i, err := f.FindInventoryItemById(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return Render(c, 200, components.InventoryItemForm(i.Id, strconv.Itoa(i.Amount)))
	}
}

func PutInventory(finder putInventoryDeps) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		amount := c.FormValue("amount")

		a, err := strconv.Atoi(amount)
		if err != nil { // TODO control error and map to http
			return err
		}

		input := inventorymanager.UpdateInventoryItemInput{
			Id:     id,
			Amount: a,
		}
		err = finder.UpdateInventoryItem(c.Request().Context(), input)
		if err != nil { // TODO control error and map to http
			return err
		}

		i, err := finder.FindInventoryItemById(c.Request().Context(), id)
		if err != nil { // TODO control error and map to http
			return err
		}

		return Render(c, 200, components.InventoryValue(i.Id, strconv.Itoa(i.Amount)))

	}
}
