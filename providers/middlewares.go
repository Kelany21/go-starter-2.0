package providers

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/interfaces/middleware"
)

func middlewares(r *gin.Engine) *gin.Engine {
	/// run cors middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Language())
	r.Use(middleware.Jsonapi())

	return r
}
