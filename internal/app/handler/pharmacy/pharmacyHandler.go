package pharmacy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	pharmacyService "github.com/tadilbek11kz/ePharma-backend/internal/service/pharmacy"
	"github.com/tadilbek11kz/ePharma-backend/internal/util"
	"github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
)

type Handler struct {
	Service    pharmacyService.Service
	middleware *middleware.JWTMiddleware
}

func New(service pharmacyService.Service, middleware *middleware.JWTMiddleware) *Handler {
	return &Handler{
		Service:    service,
		middleware: middleware,
	}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	pharmacyRouter := router.Group("/pharmacy")
	pharmacyRouter.Use(handler.middleware.New())
	pharmacyRouter.POST("/", handler.createPharmacy)
	pharmacyRouter.GET("/", handler.getAllPharmacies)
	pharmacyRouter.GET("/:id", handler.getPharmacy)
	pharmacyRouter.PUT("/:id", handler.updatePharmacy)
	pharmacyRouter.DELETE("/:id", handler.deletePharmacy)
}

// createPharmacy godoc
// @Summary      Create a pharmacy
// @Description  create pharmacy
// @Tags         pharmacy
// @Accept       json
// @Produce      json
// @Param        request   body      pharmacy.CreatePharmacyRequest  true  "body"
// @Success      201 {object} pharmacy.Pharmacy
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /pharmacy [post]
func (h *Handler) createPharmacy(c *gin.Context) {
	var req pharmacy.CreatePharmacyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	pharmacy, err := h.Service.CreatePharmacy(c.Request.Context(), req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create pharmacy: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusCreated, pharmacy)
}

// getAllPharmacies godoc
// @Summary      Get all pharmacies
// @Description  get all pharmacies
// @Tags         pharmacy
// @Accept       json
// @Produce      json
// @Success      200 {object} []pharmacy.Pharmacy
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /pharmacy [get]
func (h *Handler) getAllPharmacies(c *gin.Context) {
	pharmacies, err := h.Service.GetAllPharmacies(c.Request.Context())

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.Respond(c, http.StatusOK, pharmacies)
}

// getPharmacy godoc
// @Summary      Get a pharmacy
// @Description  get pharmacy
// @Tags         pharmacy
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "Pharmacy ID"
// @Success      200 {object} pharmacy.Pharmacy
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /pharmacy/{id} [get]
func (h *Handler) getPharmacy(c *gin.Context) {
	id := c.Param("id")

	pharmacy, err := h.Service.GetPharmacy(c.Request.Context(), id)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.Respond(c, http.StatusOK, pharmacy)
}

// updatePharmacy godoc
// @Summary      Update a pharmacy
// @Description  update pharmacy
// @Tags         pharmacy
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "Pharmacy ID"
// @Param        request   body      pharmacy.UpdatePharmacyRequest  true  "body"
// @Success      200 {object} pharmacy.Pharmacy
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /pharmacy/{id} [put]
func (h *Handler) updatePharmacy(c *gin.Context) {
	id := c.Param("id")

	var req pharmacy.UpdatePharmacyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	pharmacy, err := h.Service.UpdatePharmacy(c.Request.Context(), id, req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to update pharmacy: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, pharmacy)
}

// deletePharmacy godoc
// @Summary      Delete a pharmacy
// @Description  delete pharmacy
// @Tags         pharmacy
// @Accept       json
// @Produce      json
// @Param        id     path      string  true  "Pharmacy ID"
// @Success      204
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /pharmacy/{id} [delete]
func (h *Handler) deletePharmacy(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeletePharmacy(c.Request.Context(), id)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete pharmacy: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusNoContent, nil)
}
