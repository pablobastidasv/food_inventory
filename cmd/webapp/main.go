package main

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablobastidasv/fridge_inventory/handler"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/server"
	"github.com/pablobastidasv/fridge_inventory/storage/postgres"
)

func main() {
	e := echo.New()

	e.Static("/statics", "assets")

    
	// Middleware
	e.Use(middleware.Logger())
    e.Use(middleware.Recover())


	connStr := "host=localhost user=postgres dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	store := postgres.New(db)
	manager := inventorymanager.New(store)

	e.GET("/products/new", handler.GetProductsNew())
	e.POST("/products", handler.PostProducts(manager), server.WithTransaction)

	e.Logger.Fatal(e.Start(":1323"))

}
