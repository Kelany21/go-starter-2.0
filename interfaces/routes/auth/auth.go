package auth

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/app"
	"golang-ddd-starter/interfaces"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	interfaces.POST(r, app.AuthApplication.Login)
	interfaces.POST(r, app.AuthApplication.Logout)
	interfaces.POST(r, app.AuthApplication.Register)
	interfaces.POST(r, app.AuthApplication.Reset)
	interfaces.POST(r, app.AuthApplication.Recover)

	return r
}
