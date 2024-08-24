package inventorymanager

import (
	"context"

	"github.com/pablobastidasv/fridge_inventory/types"
)


func (m *inventoryManager) ListInventoryItems(c context.Context) ([]types.InventoryItem, error) {
    return m.store.ListInventoryItems(c)
}
