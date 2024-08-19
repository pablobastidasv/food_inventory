package storage

import (
	"context"

	"github.com/pablobastidasv/fridge_inventory/types"
)

type Store interface {
	FindCategory(context.Context, string) (*types.Category, error)
	SaveProduct(context.Context, types.Product) error
}

