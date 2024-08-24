package inventorymanager

import (
	"context"

	"github.com/pablobastidasv/fridge_inventory/types"
)

func (m *inventoryManager) ListInventoryItems(c context.Context) ([]types.InventoryItem, error) {
	return m.store.ListInventoryItems(c)
}

func (m *inventoryManager) FindInventoryItemById(c context.Context, id string) (*types.InventoryItem, error) {
	return m.store.FindInventoryItemById(c, id)
}

func (m *inventoryManager) UpdateInventoryItem(c context.Context, input UpdateInventoryItemInput) error {
	id := input.Id
	amount := input.Amount

	return m.store.UpdateInventoryItem(c, id, amount)
}
