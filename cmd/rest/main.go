package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hifat/cost-calculator-api/config"
	"github.com/hifat/cost-calculator-api/pkg/database"
	"github.com/hifat/cost-calculator-api/pkg/initial"
	"github.com/hifat/cost-calculator-api/router"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/framework"
)

func main() {
	cfg := config.LoadAppConfig("./", ".env")
	ctx := context.Background()

	dbClient := database.MongoConnect(ctx, &cfg.Db)
	defer dbClient.Disconnect(ctx)

	db := dbClient.Database(cfg.Db.Name)

	app := fiber.New()
	engine := framework.NewFiberEngineCtx(fiber.New())

	app.Use(cors.New())

	engine.Use(func(c core.IHttpCtx) {
		// Set CORS headers
		c.SetHeader("Access-Control-Allow-Origin", "*")
		c.SetHeader("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.SetHeader("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.SetHeader("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Method() == "OPTIONS" {
			// c.Status(204)
			return
		}

		c.Next()
	})

	r := router.New(engine, cfg, db, initial.Logger)

	engine.Get("/health", func(ic core.IHttpCtx) {
		ic.JSON(200, map[string]any{
			"message": "ok",
		})
	})

	r.InventoryRouter()
	r.RecipeRouter()

	engine.Listener(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
