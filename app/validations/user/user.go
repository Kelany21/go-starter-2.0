package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

/**
* validate store user request
 */
func Store(g *gin.Context, request *models.User) bool {
	lang := helpers.GetCurrentLang(g)
	/// Validation rules
	rules := govalidator.MapData{
		"name":     []string{"required", "min:6", "max:50"},
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"required", "min:6", "max:50"},
		"block":    []string{"required", "between:1,2"},
		"role":     []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"name":     []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"email":    []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50"), helpers.Email(lang)},
		"password": []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"block":    []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
		"role":     []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
	}

	opts := govalidator.Options{
		Request:         g.Request,     // request object
		Rules:           rules, // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}

/**
* validate update user request
 */
func Update(g *gin.Context, request *models.User) bool {
	lang := helpers.GetCurrentLang(g)
	/// Validation rules
	rules := govalidator.MapData{
		"name":     []string{"required", "min:6", "max:50"},
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"max:50"},
		"block":    []string{"required", "between:1,2"},
		"role":     []string{"required", "between:1,2"},
	}

	messages := govalidator.MapData{
		"name":     []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50")},
		"email":    []string{helpers.Required(lang), helpers.Min(lang, "6"), helpers.Max(lang, "50"), helpers.Email(lang)},
		"password": []string{helpers.Max(lang, "50")},
		"block":    []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
		"role":     []string{helpers.Required(lang), helpers.Min(lang, "1"), helpers.Max(lang, "2")},
	}

	opts := govalidator.Options{
		Request:         g.Request,     // request object
		Rules:           rules, // rules map
		Data:            request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	return !helpers.ReturnNotValidRequest(govalidator.New(opts), g)
}
