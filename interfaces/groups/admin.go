package groups

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/interfaces/routes/users"
)

/***
* any route here will add after /admin
* admin only  will have access this groups
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	users.Routes(r)

	return r
}
