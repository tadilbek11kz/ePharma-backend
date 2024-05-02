package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	productService "github.com/tadilbek11kz/ePharma-backend/internal/service/product"
	"github.com/tadilbek11kz/ePharma-backend/internal/util"
	"github.com/tadilbek11kz/ePharma-backend/pkg/product"
)

type Handler struct {
	Service    productService.Service
	middleware *middleware.JWTMiddleware
}

func New(service productService.Service, middleware *middleware.JWTMiddleware) *Handler {
	return &Handler{
		Service:    service,
		middleware: middleware,
	}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	productRouter := router.Group("/product")
	productRouter.Use(handler.middleware.New())
	productRouter.POST("/", handler.createProduct)
	productRouter.GET("/", handler.getAllProducts)
	productRouter.GET("/:id", handler.getProduct)
	productRouter.PUT("/:id", handler.updateProduct)
	productRouter.DELETE("/:id", handler.deleteProduct)
}

// createProduct godoc
// @Summary      Create a product
// @Description  create product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request   body      product.CreateProductRequest  true  "body"
// @Success      201 {object} product.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product [post]
func (h *Handler) createProduct(c *gin.Context) {
	var req product.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	product, err := h.Service.CreateProduct(c.Request.Context(), req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create product: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusCreated, product)
}

// getAllProducts godoc
// @Summary      Get all products
// @Description  get products
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200 {object} []product.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.Service.GetAllProducts(c.Request.Context())
	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get products: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, products)
}

// getProduct godoc
// @Summary      Get a product
// @Description  get product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id     path    string     true        "Product ID"
// @Success      200 {object} product.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product/{id} [get]
func (h *Handler) getProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.Service.GetProduct(c.Request.Context(), id)
	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to get product: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, product)
}

// updateProduct godoc
// @Summary      Update a product
// @Description  update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id     path    string     true        "Product ID"
// @Param        request   body      product.UpdateProductRequest  true  "body"
// @Success      200 {object} product.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	id := c.Param("id")

	var req product.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	product, err := h.Service.UpdateProduct(c.Request.Context(), id, req)

	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to update product: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusOK, product)
}

// deleteProduct godoc
// @Summary      Delete a product
// @Description  delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id     path    string     true        "Product ID"
// @Success      204
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete product: " + err.Error()})
		return
	}

	util.Respond(c, http.StatusNoContent, nil)
}
