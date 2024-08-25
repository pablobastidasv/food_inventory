package main

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablobastidasv/fridge_inventory/handler"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/server"
	"github.com/pablobastidasv/fridge_inventory/storage/postgres"
	slogecho "github.com/samber/slog-echo"
)

func main() {
	_ = godotenv.Load()

	// creating custom log
	logger := newLogger()

	// Loading database
	db := newDb()
	defer db.Close()

	e := echo.New()

	e.Static("/statics", "assets")

	// Middlewares
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	store := postgres.New(db)
	manager := inventorymanager.New(store)

	e.GET("/", handler.GetMainIndex())

	invGroup := e.Group("/inventories")
	invGroup.GET("", handler.GetInventoryItems(manager))
	invGroup.PUT("/:id", handler.PutInventory(manager))
	invGroup.GET("/:id/edit", handler.GetInventoryForm(manager))

	prdGroup := e.Group("/products")
	prdGroup.GET("", handler.GetProducts(manager))
	prdGroup.POST("", handler.PostProducts(manager), server.WithTransaction)
	prdGroup.GET("/new", handler.GetProductsForm(manager))
	prdGroup.DELETE("/:id", handler.DeleteProduct(manager), server.WithTransaction)

	e.Logger.Fatal(e.Start(":8080"))
}

func newLogger() *slog.Logger {
	var logLevel slog.Level

	switch os.Getenv("ENV") {
	case "DEV":
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	}))

	slog.SetDefault(logger)

	return logger

}

func newDb() *sql.DB {
	connStr := os.Getenv("DBSTRING")
    if connStr == "" {
        panic("DBSTRING environment variable must be provided to initialize the server")
    }

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
