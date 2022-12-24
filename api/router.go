package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func R(r *gin.Engine) {
	v1 := r.Group("/" + "v0.1")
	{
		v1.GET("/test", r_test)
		v1.GET("/parser/:token", r_checkToken)
		v1.GET("/parser/:token/:b64url", r_profile)
		v1.GET("/parser/:token/:b64url/:os", r_parser)
	}

}

func r_test(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
