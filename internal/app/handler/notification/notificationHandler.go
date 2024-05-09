package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	notificationService "github.com/tadilbek11kz/ePharma-backend/internal/service/notification"
	"github.com/tadilbek11kz/ePharma-backend/internal/util"
	"github.com/tadilbek11kz/ePharma-backend/pkg/notification"
)

type Handler struct {
	Service    notificationService.Service
	middleware *middleware.JWTMiddleware
}

func New(service notificationService.Service, middleware *middleware.JWTMiddleware) *Handler {
	return &Handler{
		Service:    service,
		middleware: middleware,
	}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	notificationRouter := router.Group("/notification")
	// notificationRouter.Use(handler.middleware.New())
	notificationRouter.POST("/", handler.createNotification)
	notificationRouter.DELETE("/:id", handler.deleteNotification)
}

// createNotification godoc
// @Summary Create notification
// @Description Create notification
// @Tags notification
// @Accept json
// @Produce json
// @Param request body notification.CreateNotificationRequest true "body"
// @Success 201 {object} notification.Notification
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /notification [post]
func (h *Handler) createNotification(c *gin.Context) {
	var req notification.CreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	notification, err := h.Service.CreateNotification(c.Request.Context(), req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create notification: " + err.Error()})
		return
	}

	util.Respond(c, 201, notification)
}

// deleteNotification godoc
// @Summary Delete notification
// @Description Delete notification
// @Tags notification
// @Accept json
// @Produce json
// @Param id path string true "Notification ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /notification/{id} [delete]
func (h *Handler) deleteNotification(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteNotification(c.Request.Context(), id)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete notification: " + err.Error()})
		return
	}

	util.Respond(c, 204, nil)
}
