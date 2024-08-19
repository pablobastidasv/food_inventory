package inventorymanager

import (
	"context"
	"log/slog"

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
)

type InventoryManager struct {
	store storage.Store
}

func New(store storage.Store) *InventoryManager {
	return &InventoryManager{
		store: store,
	}
}

type CreateProductInput struct {
	Id       string
	Name     string
	Category string
}

func (m *InventoryManager) CreateProduct(c context.Context, input CreateProductInput) (types.Product, error) {
	slog.Debug("Create product Use Case")

	categoryCode := input.Category
	category, err := m.store.FindCategory(c, categoryCode)
	if err != nil {
		return types.Product{}, err
	}

	product := types.Product{
		Id:       input.Id,
		Name:     input.Name,
		Category: *category,
	}
	err = m.store.SaveProduct(c, product)
	if err != nil {
		return types.Product{}, err
	}

	return product, nil
}

func (m *InventoryManager) ListProduct(c context.Context) ([]types.Product, error){
    return m.store.ListProducts(c)
}
