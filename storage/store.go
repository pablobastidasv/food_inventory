package storage

import (
	"context"

	"github.com/pablobastidasv/fridge_inventory/types"
)

type Store interface {
	SaveProduct(context.Context, types.Product) error
	ListProducts(context.Context) ([]types.Product, error)
	DeleteProduct(context.Context, string) error

	FindCategory(context.Context, string) (*types.Category, error)
	ListCategories(context.Context) ([]types.Category, error)

	ListInventoryItems(context.Context) ([]types.InventoryItem, error)
	FindInventoryItemById(context.Context, string) (*types.InventoryItem, error)
	UpdateInventoryItem(context.Context, string, int) error
}
