package main

import (
	"log/slog"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablobastidasv/fridge_inventory/auth"
	"github.com/pablobastidasv/fridge_inventory/db"
	"github.com/pablobastidasv/fridge_inventory/handler"
	"github.com/pablobastidasv/fridge_inventory/internals/inventorymanager"
	"github.com/pablobastidasv/fridge_inventory/server"
	"github.com/pablobastidasv/fridge_inventory/storage/postgres"
	slogecho "github.com/samber/slog-echo"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	// creating custom log
	logger := newLogger()

	// Loading database
	dbConn := db.NewPostgresDb()
	defer dbConn.Close()

	e := echo.New()

	e.Static("/statics", "assets")

	// Middlewares
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret")))) // TODO: real secret value here

	store := postgres.New(dbConn)
	manager := inventorymanager.New(store)


	// ======= Begin: Authentication
    ah := auth.New("/auth/login")
    g := e.Group("/auth")
	g.GET("/login", ah.GetLogin)
	g.GET("/callback", ah.GetCallback)
	g.GET("/logout", ah.GetLogout)
	// ======= End: Authentication

	e.GET("/", handler.GetMainIndex(), ah.PageMiddleware)

	invGroup := e.Group("/inventories", ah.FragmentMiddleware)
	invGroup.GET("", handler.GetInventoryItems(manager))
	invGroup.PUT("/:id", handler.PutInventory(manager))
	invGroup.GET("/:id/edit", handler.GetInventoryForm(manager))

	prdGroup := e.Group("/products")
	prdGroup.GET("", handler.GetProducts(manager), ah.PageMiddleware)
	prdGroup.POST("", handler.PostProducts(manager), server.WithTransaction)
	prdGroup.GET("/new", handler.GetProductsForm(manager))
	prdGroup.DELETE("/:id", handler.DeleteProduct(manager), server.WithTransaction)

	e.Logger.Fatal(e.Start(":8080"))
}

func newLogger() *slog.Logger {
	var logger *slog.Logger

	switch os.Getenv("ENV") {
	case "DEV":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))

	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: true,
		}))

	}

	slog.SetDefault(logger)

	return logger

}
