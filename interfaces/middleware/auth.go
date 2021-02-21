package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/helpers"
	"golang-ddd-starter/infrastructure"
	"strings"
)

/**
* This middleware will Allow only auth user
* Will block not auth user
* if user allow to access then this middleware will add
* one header with user information you can use later (AUTH_DATA)
* in function you call
 */
func Auth() gin.HandlerFunc {
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
		user, _ := userRepo.Get("token = ?", adminToken)
		if user.UUID.String() == "" {
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
		g.Request.Header.Set("AUTH_DATA", string(userJson))
		g.Next()
	}
}
