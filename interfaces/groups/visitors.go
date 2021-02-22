package groups

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/interfaces/routes/auth"
)

/***
* any route here will add after /
* anyone will have access this groups
 */
func Visitor(r *gin.RouterGroup) *gin.RouterGroup {
	auth.Routes(r)

	return r
}
