package users

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/app"
	"golang-ddd-starter/interfaces"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	interfaces.GET(r, app.UserApplication.List)
	interfaces.POST(r, app.UserApplication.Create)
	interfaces.PUT(r, app.UserApplication.Update, "uuid")
	interfaces.GET(r, app.UserApplication.Show, "uuid")
	interfaces.DELETE(r, app.UserApplication.Delete, "uuid")

	return r
}
