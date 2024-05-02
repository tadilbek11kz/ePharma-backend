package appserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/config"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/handler/auth"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/handler/pharmacy"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/handler/product"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/handler/swagger"
	"github.com/tadilbek11kz/ePharma-backend/internal/connections"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	"github.com/tadilbek11kz/ePharma-backend/internal/util/logger"

	"go.uber.org/fx"
)

func RegisterHooks(
	lifecycle fx.Lifecycle,
	config *config.TomlConfig,
	router *gin.Engine,
	logger logger.Logger,
	conns *connections.Connections,
	cors *middleware.CORSMiddleware,
	authHandler *auth.Handler,
	pharmacyHandler *pharmacy.Handler,
	productHandler *product.Handler,
	swaggerHandler *swagger.Handler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				router.Use(cors.New())

				auth.RegisterRoutes(router, authHandler)
				pharmacy.RegisterRoutes(router, pharmacyHandler)
				product.RegisterRoutes(router, productHandler)
				swagger.RegisterRoutes(router, swaggerHandler)

				go func() {
					err := http.ListenAndServe(config.BindAddr, router)
					if err != nil {
						logger.Info(err.Error(), map[string]interface{}{"BindAddr": config.BindAddr}) //nolint
					}

					logger.Info(fmt.Sprintf("Listening on %s", config.BindAddr), map[string]interface{}{})
				}()

				return nil
			},
			OnStop: func(context.Context) error {
				conns.Close()
				logger.Info("Server stopped", map[string]interface{}{})
				return nil
			},
		},
	)
}
