package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/http-swagger"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/appserver"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/config"
	authHandler "github.com/tadilbek11kz/ePharma-backend/internal/app/handler/auth"
	inventoryHandler "github.com/tadilbek11kz/ePharma-backend/internal/app/handler/inventory"
	notificationHandler "github.com/tadilbek11kz/ePharma-backend/internal/app/handler/notification"
	pharmacyHandler "github.com/tadilbek11kz/ePharma-backend/internal/app/handler/pharmacy"
	productHandler "github.com/tadilbek11kz/ePharma-backend/internal/app/handler/product"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/handler/swagger"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	"github.com/tadilbek11kz/ePharma-backend/internal/connections"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	"github.com/tadilbek11kz/ePharma-backend/internal/service/auth"
	"github.com/tadilbek11kz/ePharma-backend/internal/service/inventory"
	"github.com/tadilbek11kz/ePharma-backend/internal/service/notification"
	"github.com/tadilbek11kz/ePharma-backend/internal/service/pharmacy"
	"github.com/tadilbek11kz/ePharma-backend/internal/service/product"
	logrus_log "github.com/tadilbek11kz/ePharma-backend/internal/util/logger/logrus-log"
	"go.uber.org/fx"
)

// @title ePharma API
// @version 1.0
// @description This is a service for ePharma project
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 10.225.78.198:8080
// @BasePath /
func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	fx.New(
		fx.Provide(
			config.NewConfig,
			logrus_log.New,
			gin.Default,
			connections.New,
			store.New,
			middleware.NewCORSMiddleware,
			middleware.NewJWTMiddleware,
			auth.New,
			authHandler.New,
			pharmacy.New,
			pharmacyHandler.New,
			product.New,
			productHandler.New,
			inventory.New,
			inventoryHandler.New,
			notification.New,
			notificationHandler.New,
			swagger.New,
		),

		fx.Invoke(
			appserver.RegisterHooks,
		),
	).Run()
}
