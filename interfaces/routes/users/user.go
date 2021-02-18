package users

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/app"
	"golang-ddd-starter/interfaces"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	interfaces.GET(r, app.UserApplication.List)
	interfaces.GET(r, app.UserApplication.Paginate)
	interfaces.POST(r, app.UserApplication.Create)
	interfaces.PUT(r, app.UserApplication.Update, "id")
	interfaces.GET(r, app.UserApplication.Show, "id")
	interfaces.DELETE(r, app.UserApplication.Delete, "id")

	return r
}