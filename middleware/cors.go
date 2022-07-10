package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Cache-Control", "max-age=36000")

		m := map[string]struct{}{
			"GET": {}, "POST": {}, "PUT": {}, "DELETE": {},
		}
		if _, ok := m[c.Request.Method]; !ok {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
