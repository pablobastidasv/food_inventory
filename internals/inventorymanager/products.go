package inventorymanager

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/pablobastidasv/fridge_inventory/types"
)

var ErrProductHasStock = errors.New("product has stock")

func (m *inventoryManager) CreateProduct(c context.Context, input CreateProductInput) (types.Product, error) {
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

	item := types.InventoryItem{
		Id:      uuid.NewString(),
		Product: product,
		Amount:  0,
	}

	err = m.store.CreateInventoryItem(c, item)
	if err != nil {
		return types.Product{}, err
	}

	return product, nil
}

func (m *inventoryManager) ListProduct(c context.Context) ([]types.Product, error) {
	return m.store.ListProducts(c)
}

func (m *inventoryManager) ListCategories(c context.Context) ([]types.Category, error) {
	return m.store.ListCategories(c)
}

func (m *inventoryManager) DeleteProduct(ctx context.Context, productId string) error {
    i, err := m.store.FindInventoryItemByProduct(ctx, types.Product{Id: productId})
    if err != nil {
        return err
    }
    if i != nil && i.HasStock() {
       return ErrProductHasStock
    }
	return m.store.DeleteProduct(ctx, productId)
}
