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

type MenuRoutes struct {
	log logger.LoggerInterface
	uc  usecase.Menuer
}

func newMenuRoutes(handler *gin.RouterGroup, log logger.LoggerInterface, db *pg.DB) {
	menuRepo := repo.NewMenuRepo(db, log)
	uc := usecase.NewMenuUsecase(log, menuRepo)
	route := MenuRoutes{log: log, uc: uc}

	h := handler.Group("/menu")
	{
		h.GET("/", route.GetMenus)
	}
}

// @Summary     Get All Menu's
// @Tags  	    Menu
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param		req query model.MenuRequest true "Menu Request"
// @Success     200 {object} model.MenuResponse
// @Router      /v1/menu [get]
func (r *MenuRoutes) GetMenus(c *gin.Context) {
	req := new(model.MenuRequest)
	if err := c.ShouldBind(req); err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusBadRequest, err, c))
		return
	}

	ctx := c.Request.Context()
	mm, err := r.uc.GetMenues(ctx, req)
	if err != nil {
		r.log.Error(error_helper.AbortOnError(http.StatusInternalServerError, err, c))
		return
	}

	resp := &model.Response{
		Message:   "Success get menues",
		Status:    http.StatusText(http.StatusOK),
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      mm,
	}
	c.JSON(http.StatusOK, resp)
}
