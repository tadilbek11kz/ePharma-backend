package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	inventoryService "github.com/tadilbek11kz/ePharma-backend/internal/service/inventory"
	"github.com/tadilbek11kz/ePharma-backend/internal/util"
	"github.com/tadilbek11kz/ePharma-backend/pkg/inventory"
)

type Handler struct {
	Service    inventoryService.Service
	middleware *middleware.JWTMiddleware
}

func New(service inventoryService.Service, middleware *middleware.JWTMiddleware) *Handler {
	return &Handler{
		Service:    service,
		middleware: middleware,
	}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	inventoryRouter := router.Group("/inventory")
	inventoryRouter.Use(handler.middleware.New())
	inventoryRouter.POST("/", handler.createInventory)
	inventoryRouter.GET("/", handler.getAllInventories)
	// inventoryRouter.GET("/:id", handler.getInventory)
	// inventoryRouter.PUT("/:id", handler.updateInventory)
	// inventoryRouter.DELETE("/:id", handler.deleteInventory)
}

// createInventory godoc
// @Summary Create inventory
// @Description Create inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param request body inventory.CreateInventoryRequest true "body"
// @Success 201 {object} inventory.Inventory
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /inventory [post]
func (h *Handler) createInventory(c *gin.Context) {
	var req inventory.CreateInventoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	inventory, err := h.Service.CreateInventory(c.Request.Context(), req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create inventory: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusCreated, inventory)
}

// getAllInventories godoc
// @Summary Get all inventories
// @Description Get all inventories
// @Tags inventory
// @Accept json
// @Produce json
// @Success 200 {object} []inventory.Inventory
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /inventory [get]
func (h *Handler) getAllInventories(c *gin.Context) {
	inventories, err := h.Service.GetAllInventories(c.Request.Context())

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get all inventories: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, inventories)
}

// getInventory godoc
// @Summary Get inventory
// @Description Get inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 200 {object} inventory.Inventory
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /inventory/{id} [get]
func (h *Handler) getInventory(c *gin.Context) {
	id := c.Param("id")

	inventory, err := h.Service.GetInventory(c.Request.Context(), id)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get inventory: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, inventory)
}

// updateInventory godoc
// @Summary Update inventory
// @Description Update inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Param request body inventory.UpdateInventoryRequest true "body"
// @Success 200 {object} inventory.Inventory
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /inventory/{id} [put]
func (h *Handler) updateInventory(c *gin.Context) {
	id := c.Param("id")

	var req inventory.UpdateInventoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	inventory, err := h.Service.UpdateInventory(c.Request.Context(), id, req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to update inventory: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, inventory)
}

// deleteInventory godoc
// @Summary Delete inventory
// @Description Delete inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /inventory/{id} [delete]
func (h *Handler) deleteInventory(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteInventory(c.Request.Context(), id)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete inventory: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusNoContent, nil)
}
