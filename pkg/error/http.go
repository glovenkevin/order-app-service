package error

import (
	"context"
	"net/http"
	"order-app/domain/model"
	"time"

	"github.com/gin-gonic/gin"
)

func SendError(code int, err error, g *gin.Context) string {
	r := model.Response{
		Message:   err.Error(),
		Status:    getStatusDesc(code),
		Timestamp: time.Now().Format(time.RFC3339),
	}
	g.JSON(http.StatusInternalServerError, r)

	return err.Error()
}

func AbortOnError(code int, err error, g *gin.Context) string {
	r := model.Response{
		Message:   err.Error(),
		Status:    getStatusDesc(code),
		Timestamp: time.Now().Format(time.RFC3339),
	}
	g.AbortWithStatusJSON(code, r)

	return err.Error()
}

func AbortAuthenticated(g *gin.Context) string {
	return AbortOnError(http.StatusUnauthorized, ErrReqNotAuthenticated, g)
}

func ContextError(c context.Context) error {
	switch c.Err() {
	case context.Canceled:
		return ErrRequestCanceled
	case context.DeadlineExceeded:
		return ErrDeadlineExceeded
	default:
		return nil
	}
}
