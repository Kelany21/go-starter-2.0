package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/helpers"
	"golang-ddd-starter/infrastructure"
	"strings"
)

/**
* This middle ware will Allow only
* Will block not admin role admin role is (2)
* if user allow to access then this middleware will add
* one header with user information you can use later (ADMIN_DATA)
* in function you call
 */
func Admin() gin.HandlerFunc {
	return func(g *gin.Context) {
		/// get Authorization header to check if user send it first
		adminToken := strings.ReplaceAll(g.GetHeader("Authorization"), "Bearer ", "")
		if adminToken == "" {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// check if token exits in database
		userRepo := infrastructure.NewUserRepository(infrastructure.DB)
		user, _ := userRepo.Get("token = ? and role = ?", adminToken, 2)
		if user.UUID.String() == "00000000-0000-0000-0000-000000000000" {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// check if user block or not
		if user.Block != 2 {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// not set header with user information
		userJson, _ := json.Marshal(&user)
		g.Request.Header.Set("ADMIN_DATA", string(userJson))
		g.Next()
	}
}
