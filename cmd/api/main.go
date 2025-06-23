package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	core "github.com/hifat/goroger-core"
	"github.com/hifat/goroger-core/framework"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/constant"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/initial"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"github.com/hifat/mallow-sale-api/router"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/fiber-swagger/example/docs"
)

// @title           Mallow Sale API
// @version         1.0
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("please give service name and env path: go run . <env_path>")
	}

	cfg := config.LoadAppConfig(args[1])
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
		c.SetHeader("Access-Control-Allow-Methods", strings.Join([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		}, ","))
		c.SetHeader("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Api-Key")
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

	grpc, err := rpc.NewGRPCClient(cfg)
	if err != nil {
		panic(err)
	}
	defer grpc.CloseAll()

	r := router.New(engine, cfg, db, initial.Logger, validateRegis, grpc)

	engine.Get("/health", func(ic core.IHttpCtx) {
		ic.JSON(200, map[string]any{
			"message": "ok",
		})
	})

	switch cfg.App.Service {
	case constant.ServiceInventory:
		r.InventoryRouter()
	case constant.ServiceRecipe:
		r.RecipeRouter()
	case constant.ServiceUsageUnit:
		r.UsageUnitRouter()
	}

	engine.Listener(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
