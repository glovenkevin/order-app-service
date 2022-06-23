package middleware

import (
	"fmt"
	"net/http"
	"order-app/domain/model"
	"order-app/pkg/logger"
	"order-app/pkg/time"

	"github.com/gin-gonic/gin"
)

func PanicRecovery(l logger.LoggerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()

			if err != nil {
				l.Error(err)
				resp := &model.Response{
					Message:   fmt.Sprintf("[ERROR] Cause: %s", err.(error).Error()),
					Status:    http.StatusText(http.StatusInternalServerError),
					Timestamp: time.GetNow(),
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
				return
			}
		}()
		c.Next()
	}
}
