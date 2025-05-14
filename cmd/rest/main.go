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
	"github.com/hifat/goroger-core/rules"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title           Mallow Sale API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
func main() {
	cfg := config.LoadAppConfig("./", ".env")
	ctx := context.Background()

	dbClient := database.MongoConnect(ctx, &cfg.Db)
	defer dbClient.Disconnect(ctx)

	db := dbClient.Database(cfg.Db.Name)

	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

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

	validateRegis, err := rules.Register()
	if err != nil {
		panic(err)
	}

	r := router.New(engine, cfg, db, initial.Logger, validateRegis)

	engine.Get("/health", func(ic core.IHttpCtx) {
		ic.JSON(200, map[string]any{
			"message": "ok",
		})
	})

	r.InventoryRouter()
	r.RecipeRouter()

	engine.Listener(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
