package v1

import (
	"net/http"
	"order-app/domain/model"
	"order-app/domain/usecase"
	"order-app/domain/usecase/repo"
	"order-app/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type AuthRoutes struct {
	log logger.ILogger
	uc  usecase.Auther
}

func newAuthRoutes(handler *gin.RouterGroup, log logger.ILogger, db *pg.DB) {
	userRepo := repo.NewUserRepo(log, db)
	roleRepo := repo.NewRoleRepo(log, db)

	uc := usecase.NewAuthUseCase(log, userRepo, roleRepo)

	route := AuthRoutes{log: log, uc: uc}

	h := handler.Group("/auth")
	{
		h.POST("/login", route.login)
		h.POST("/register", route.register)
	}
}

// @Summary     User login
// @Description Authenticate user whether it is valid or not
// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param		req body model.LoginRequest true "Login request"
// @Success     200 {object} model.LoginResponse
// @Failure		400 {object} v1.ErrorResponse
// @Failure     500 {object} v1.ErrorResponse
// @Router      /api/v1/auth/login [get]
func (r *AuthRoutes) login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error(err)
		errorResponse(c, http.StatusBadRequest, "bad request")

		return
	}

	resp := &model.LoginResponse{
		Token:   "",
		Message: "Test masuk",
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary     Register new user
// @Description Registering new user
// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param		req body model.RegisterRequest true "Register request"
// @Success     200 {object} model.SuccessResponse
// @Failure		400 {object} v1.ErrorResponse
// @Failure     500 {object} v1.ErrorResponse
// @Router      /api/v1/auth/register [get]
func (r *AuthRoutes) register(c *gin.Context) {
	select {
	case <-c.Done():
		if c.Err() != nil {
			r.log.Errorf("c.Done(): %v", c.Err())
			errorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		return
	default:
	}

	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error(err)
		errorResponse(c, http.StatusBadRequest, "bad request")
		return
	}

	ctx := c.Request.Context()
	err := r.uc.Register(ctx, req.ToEntity())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := &model.SuccessResponse{
		Message: "Berhasil Daftar",
	}
	c.JSON(http.StatusOK, resp)
}
