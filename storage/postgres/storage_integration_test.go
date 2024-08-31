//go:build integration

package postgres_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/pablobastidasv/fridge_inventory/storage"
	"github.com/pablobastidasv/fridge_inventory/storage/postgres"
	"github.com/pablobastidasv/fridge_inventory/testutils"
	"github.com/pablobastidasv/fridge_inventory/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinygg/gofaker"
)

func TestMain(m *testing.M) {
	testutils.LoadEnv()

	postgresDb := testutils.DbInstance()
	code := m.Run()
	postgresDb.Close()

	os.Exit(code)
}

func Test_FindCategory(tt *testing.T) {
	tt.Parallel()

	var sut storage.Store
	postgresDb := testutils.DbInstance()

	sut = postgres.New(postgresDb)

	tt.Run("given a category code, when the category exists and does not have a parent, then category information is return and parent is nil", func(t *testing.T) {
		categoryCode := "VEGETABLES"

		category, err := sut.FindCategory(context.Background(), categoryCode)
		require.NoError(t, err)

		assert.Equal(t, categoryCode, category.Code)
		assert.Equal(t, "Vegetales", category.Name)
		assert.Nil(t, category.Parent)
	})

	tt.Run("given a category code, when the category does not exists, then the return is nil", func(t *testing.T) {
		categoryCode := "UNEXPECTED"

		result, err := sut.FindCategory(context.Background(), categoryCode)
		assert.NoError(t, err)

		assert.Nil(t, result)
	})

	tt.Run("given a category code, when the category contains a parent, then the parent information is returned", func(t *testing.T) {
		categoryCode := "PORK"

		category, err := sut.FindCategory(context.Background(), categoryCode)
		require.NoError(t, err)

		assert.Equal(t, categoryCode, category.Code)
		assert.Equal(t, "Cerdo", category.Name)
		assert.Equal(t, "MEAT", category.Parent.Code)
		assert.Equal(t, "Carne", category.Parent.Name)
		assert.Nil(t, category.Parent.Parent)
	})
}

func Test_SaveProduct(tt *testing.T) {
	var sut storage.Store
	postgresDb := testutils.DbInstance()

	sut = postgres.New(postgresDb)

	tt.Run("given a product, when it's saved correctly, then error is NOT returned", func(t *testing.T) {
		category, err := sut.FindCategory(context.Background(), "FRUITS")
		if err != nil {
			t.Fatal(err)
		}

		product := types.Product{
			Id:       gofaker.UUID(),
			Name:     gofaker.Fruit(),
			Category: *category,
		}

		err = sut.SaveProduct(context.Background(), product)
		assert.NoError(t, err)

		query := "select * from products where id = $1"
		_, err = postgresDb.QueryContext(context.Background(), query, product.Id)
		assert.NoError(t, err)
	})
}

func Test_ListProducts(tt *testing.T) {
	var sut storage.Store
	postgresDb := testutils.DbInstance()

	sut = postgres.New(postgresDb)

	tt.Run("given a list products request, when it goes as expected, then not empty list is returned", func(t *testing.T) {
		products, err := sut.ListProducts(context.Background())
		assert.NoError(t, err)

		assert.NotEmpty(t, products)
	})
}

func Test_ListCategories(tt *testing.T) {
	var sut storage.Store
	postgresDb := testutils.DbInstance()

	sut = postgres.New(postgresDb)

	tt.Run("given a list categories request, when it goes as expected, then not empty list is returned", func(t *testing.T) {
		categories, err := sut.ListCategories(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, categories)
	})
}

func Test_DeleteProduct(tt *testing.T) {
	db := testutils.DbInstance()
	sut := postgres.New(db)

	tt.Run("given a created product, when delete the product, then product is not available anymore", func(t *testing.T) {
		category, err := sut.FindCategory(context.Background(), "VEGETABLES")
		assert.NoError(t, err)
		product := types.Product{
			Id:       gofaker.UUID(),
			Name:     gofaker.Vegetable(),
			Category: *category,
		}
		sut.SaveProduct(context.Background(), product)

		err = sut.DeleteProduct(context.Background(), product.Id)
		assert.NoError(t, err)

		query := "select id from products where id = $1"
		r := db.QueryRowContext(context.Background(), query, product.Id)

		var id string
		err = r.Scan(&id)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}

func Test_ListInventoryItems(tt *testing.T) {
	db := testutils.DbInstance()
	sut := postgres.New(db)

	tt.Run(
		"given a list inventory items action, when list is return, then no error is return and list is not empty",
		func(t *testing.T) {
			result, err := sut.ListInventoryItems(context.Background())
			assert.NoError(t, err)

			assert.NotEmpty(t, result)
		},
	)
}

func Test_FindInventoryItemById(tt *testing.T) {
	db := testutils.DbInstance()
	sut := postgres.New(db)

	tt.Run("given an item id, when it exists, then its information is returned", func(t *testing.T) {
		itemId := "bade892a-6423-4bd7-825f-86418c80be46"

		item, err := sut.FindInventoryItemById(context.Background(), itemId)
		assert.NoError(t, err)

		assert.Equal(t, 0, item.Amount)
		assert.Equal(t, "Cerdo Wok", item.Product.Name)
		assert.Equal(t, "PORK", item.Product.Category.Code)
		assert.Equal(t, "Cerdo", item.Product.Category.Name)
	})

	tt.Run("given an item id, when it does not exists, then item is nil", func(t *testing.T) {
		itemId := "c6bdf6fc-bff1-4935-bcee-4a11900b029d"

		item, err := sut.FindInventoryItemById(context.Background(), itemId)
		assert.NoError(t, err)
		assert.Nil(t, item)
	})
}

func Test_UpdateInventoryItem(tt *testing.T) {
	db := testutils.DbInstance()
	sut := postgres.New(db)

    tt.Run("given an item to update, when it exists, then the item is updated with the new value", func(t *testing.T) {
        itemId := "fd17c6e2-ec66-4f85-a46c-c01ba24efbba"
        amount := 10

        err := sut.UpdateInventoryItem(context.Background(), itemId, amount)
        assert.NoError(t, err)

        item, err := sut.FindInventoryItemById(context.Background(), itemId)
        assert.NoError(t, err)
        assert.Equal(t, 10, item.Amount)
    })
}
