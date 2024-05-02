package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/middleware"
	authService "github.com/tadilbek11kz/ePharma-backend/internal/service/auth"
	"github.com/tadilbek11kz/ePharma-backend/internal/util"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	Service    authService.Service
	middleware *middleware.JWTMiddleware
}

func New(service authService.Service, middleware *middleware.JWTMiddleware) *Handler {
	return &Handler{
		Service:    service,
		middleware: middleware,
	}
}

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	authRouter := router.Group("/auth")
	authRouter.POST("/register", handler.registerUser)
	authRouter.POST("/login", handler.loginUser)
	authRouter.POST("/refresh", handler.refreshToken)

	authRouter.GET("/logout", handler.middleware.New(), handler.logoutUser)
}

// registerUser godoc
// @Summary      Register an user
// @Description  register user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request   body      model.CreateUserRequest  true  "body"
// @Success      201 {object} model.User
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /auth/register [post]
func (h *Handler) registerUser(c *gin.Context) {
	var credentials model.CreateUserRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		util.Respond(c, http.StatusBadRequest, "parsing err: "+err.Error())
		return
	}

	user, err := h.Service.CreateUser(c.Request.Context(), credentials)

	if err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "User with given email already exists"})
		} else {
			util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create user: " + err.Error()})
		}
		return
	}

	util.Respond(c, http.StatusCreated, user)

}

// loginUser godoc
// @Summary      Login an user
// @Description  login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request   body      model.LoginUserRequest  true  "body"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /auth/login [post]
func (h *Handler) loginUser(c *gin.Context) {
	var credentials model.LoginUserRequest

	if err := c.ShouldBindJSON(&credentials); err != nil {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Parsing err: " + err.Error()})
		return
	}

	token, err := h.Service.LoginUser(c.Request.Context(), credentials)

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) || errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
			util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "User with given credentials is not found"})
		} else {
			util.Respond(c, http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to fetch user: " + err.Error()})
		}
		return
	}

	c.SetCookie("access_token", token.AccessToken, token.AccessTokenMaxAge, "/", "localhost", false, true)
	c.SetCookie("refresh_token", token.RefreshToken, token.RefreshTokenMaxAge, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", token.AccessTokenMaxAge, "/", "localhost", false, false)

	util.Respond(c, http.StatusOK, gin.H{"status": "success", "access_token": token.AccessToken, "refresh_token": token.RefreshToken})

}

// refreshToken godoc
// @Summary      Refresh an access token
// @Description  refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /auth/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	cookie, ok := c.Params.Get("refresh")

	if !ok {
		util.Respond(c, http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid refresh token"})
		return
	}

	token, err := h.Service.RefreshToken(c.Request.Context(), cookie)
	if err != nil {
		util.Respond(c, http.StatusForbidden, gin.H{"status": "fail", "message": "Failed to refresh token: " + err.Error()})
		return
	}

	c.SetCookie("access_token", token.AccessToken, token.AccessTokenMaxAge, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", token.AccessTokenMaxAge, "/", "localhost", false, false)
}

// logoutUser godoc
// @Summary      Logout user
// @Description  logout user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /auth/logout [get]
func (h *Handler) logoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Logged out"})
}
