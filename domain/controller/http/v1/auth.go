package v1

import (
	"net/http"
	"order-app/domain/model"
	"order-app/domain/usecase"
	"order-app/domain/usecase/repo"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type AuthRoutes struct {
	log logger.LoggerInterface
	uc  usecase.Auther
}

func newAuthRoutes(handler *gin.RouterGroup, log logger.LoggerInterface, db *pg.DB) {
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
// @Success     200 {object} model.Response
// @Failure		400 {object} model.Response
// @Failure     500 {object} model.Response
// @Router      /api/v1/auth/login [post]
func (r *AuthRoutes) login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusBadRequest, err, c))
		return
	}

	ctx := c.Request.Context()
	res, err := r.uc.Login(ctx, &req)
	if err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusUnauthorized, err, c))
		return
	}

	resp := &model.Response{
		Message:   "success login",
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      res,
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary     Register new user
// @Description Registering new user
// @Tags  	    auth
// @Accept      json
// @Produce     json
// @Param		req body model.RegisterRequest true "Register request"
// @Success     200 {object} model.Response
// @Failure		400 {object} model.Response
// @Failure     500 {object} model.Response
// @Router      /api/v1/auth/register [post]
func (r *AuthRoutes) register(c *gin.Context) {
	select {
	case <-c.Done():
		if c.Err() != nil {
			r.log.Errorf("c.Done(): %v", c.Err())
			error_helper.AbortOnError(http.StatusInternalServerError, c.Err(), c)
		}
		return
	default:
	}

	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusBadRequest, err, c))
		return
	}

	ctx := c.Request.Context()
	err := r.uc.ValidateNewUser(ctx, &req)
	if err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusBadRequest, err, c))
		return
	}

	err = r.uc.Register(ctx, req.ToEntity())
	if err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusInternalServerError, err, c))
		return
	}

	resp := &model.Response{
		Message:   "success register",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, resp)
}
