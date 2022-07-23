package v1

import (
	"net/http"
	"order-app/domain/model"
	"order-app/domain/usecase"
	"order-app/domain/usecase/repo"
	error_helper "order-app/pkg/error"
	"order-app/pkg/logger"
	"order-app/pkg/time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type BannerRoutes struct {
	uc  usecase.Banner
	log logger.LoggerInterface
}

func newBannerRoutes(handler *gin.RouterGroup, log logger.LoggerInterface, db *bun.DB) {
	repo := repo.NewBannerRepo(db, log)
	uc := usecase.NewBannerUsecase(log, repo)
	route := BannerRoutes{log: log, uc: uc}

	h := handler.Group("/banners")
	{
		h.GET("/", route.GetBanners)
	}
}

// @Summary     Get Banners
// @Tags  	    Banners
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Success     200 {object} model.Response
// @Router      /v1/banners [get]
func (r *BannerRoutes) GetBanners(c *gin.Context) {
	select {
	case <-c.Done():
		if c.Err() != nil {
			r.log.Errorf("c.Done(): %v", c.Err())
			error_helper.AbortOnError(http.StatusInternalServerError, c.Err(), c)
		}
		return
	default:
	}

	ctx := c.Request.Context()
	ret, err := r.uc.GetBanners(ctx)
	if err != nil {
		r.log.Error(error_helper.SendError(http.StatusInternalServerError, err, c))
	}

	resp := &model.Response{
		Message:   "OK",
		Status:    "Success",
		Timestamp: time.GetNow(),
		Data:      ret,
	}
	c.JSON(http.StatusOK, resp)
}
