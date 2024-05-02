package swagger

import (
	_ "github.com/tadilbek11kz/ePharma-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
