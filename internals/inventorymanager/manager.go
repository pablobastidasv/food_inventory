package inventorymanager

import (
	"context"

	"github.com/pablobastidasv/fridge_inventory/storage"
	"github.com/pablobastidasv/fridge_inventory/types"
)

type (
	ProductCreator interface {
		CreateProduct(context.Context, CreateProductInput) (types.Product, error)
	}
	ProductsLister interface {
		ListProduct(context.Context) ([]types.Product, error)
	}
	CategoryLister interface {
		ListCategories(context.Context) ([]types.Category, error)
	}
	ProductDeleter interface {
		DeleteProduct(context.Context, string) error
	}

    InventoryItemsLister interface {
        ListInventoryItems(context.Context) ([]types.InventoryItem, error)
    }

	InventoryManager interface {
		ProductCreator
		ProductsLister
        ProductDeleter
		CategoryLister
        InventoryItemsLister
	}
)

type inventoryManager struct {
	store storage.Store
}

func New(store storage.Store) InventoryManager {
	return &inventoryManager{
		store: store,
	}
}

