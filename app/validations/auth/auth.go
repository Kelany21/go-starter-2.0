package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

/**
* validate store user request
 */
func Login(g *gin.Context, request *models.Login) bool { /// Validation rules
	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request:         g.Request, // request object
		Rules:           rules,     // rules map
		Data:            request,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}

/**
* validate store user request
 */
func Logout(g *gin.Context, request *models.Logout) bool { /// Validation rules
	rules := govalidator.MapData{
		"token": []string{"required", "min:6"},
	}
	opts := govalidator.Options{
		Request:         g.Request, // request object
		Rules:           rules,     // rules map
		Data:            request,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}

/**
* validate register request
 */
func Register(g *gin.Context, user *models.User) bool {
	/// Validation rules
	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:50", "email"},
		"name":     []string{"required", "min:4", "max:50"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request:         g.Request, // request object
		Rules:           rules,     // rules map
		Data:            user,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}

/**
* validate Reset request
 */
func Reset(g *gin.Context, user *models.Reset) bool {
	/// Validation rules
	rules := govalidator.MapData{
		"email": []string{"required", "min:6", "max:50", "email"},
	}
	opts := govalidator.Options{
		Request:         g.Request, // request object
		Rules:           rules,     // rules map
		Data:            user,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}

/**
* validate Recover request
 */
func Recover(g *gin.Context, user *models.Recover) bool {
	/// Validation rules
	rules := govalidator.MapData{
		"token":    []string{"required"},
		"password": []string{"required", "between:6,20"},
	}
	opts := govalidator.Options{
		Request:         g.Request, // request object
		Rules:           rules,     // rules map
		Data:            user,
		RequiredDefault: true, // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}
