package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-ddd-starter/app/transformers"
	"golang-ddd-starter/app/validations/auth"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/domian/repos"
	"golang-ddd-starter/helpers"
	"golang-ddd-starter/infrastructure"
	"os"
)

type AuthApplicationInterface interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Recover(*gin.Context)
	Reset(*gin.Context)
	Register(*gin.Context)
}

type AuthApp struct {
	repo repos.UserRepository
}

var AuthApplication *AuthApp = nil

func NewAuthApplication(db *gorm.DB) {
	AuthApplication = &AuthApp{infrastructure.NewUserRepository(db)}
}

//UserApp implements the UserApplication interface
var _ AuthApplicationInterface = &AuthApp{}

/**
* check if user have access to login in system
 */
func (a *AuthApp) Login(g *gin.Context) {
	// init user login struct to validate request
	login := new(models.Login)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if !auth.Login(g, login) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := a.repo.CheckUserExistsNotBlocked(g, login.Email, "")
	if !valid {
		return
	}
	/**
	* now check if password are valid
	* if user password is not valid we will return invalid email
	* or password
	 */
	check := helpers.CheckPasswordHash(login.Password, user.Password)
	if !check {
		helpers.ReturnNotFound(g, "your email or your password are not valid")
		return
	}
	if user.Email != "admin@admin.com" {
		// update token then return with the new data
		user.Token, _ = helpers.GenerateToken(user.Password + user.Email)
		err := a.repo.Update(&user, user.UUID)
		if err != nil {
			helpers.ReturnForbidden(g, err.Error())
			return
		}
	}
	// now user is login we can return his info
	helpers.OkResponse(g, "you are login now", transformers.UserResponse(user))
}

/**
* check if user have access to login in system
 */
func (a *AuthApp) Logout(g *gin.Context) {
	// init user login struct to validate request
	logout := new(models.Logout)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if !auth.Logout(g, logout) {
		return
	}
	user, err := a.repo.Get("token = ?", logout.Token)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	if user.Email != "admin@admin.com" {
		// update token then return with the new data
		user.Token, _ = helpers.GenerateToken(user.Password + user.Email)
		err = a.repo.Update(&user, user.UUID)
		if err != nil {
			helpers.ReturnForbidden(g, err.Error())
			return
		}
	}
	// now user is login we can return his info
	helpers.OkResponse(g, "you are logout now", transformers.UserResponse(user))
}

/**
* Register new user on system
 */
func (a *AuthApp) Register(g *gin.Context) {
	// init visitor login struct to validate request
	user := new(models.User)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if !auth.Register(g, user) {
		return
	}
	/**
	* check if this email exists database
	* if this email found will return
	 */
	anotherUser, _ := a.repo.Get("email = ?", user.Email)
	if anotherUser.UUID.String() != "00000000-0000-0000-0000-000000000000" {
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	/**
	* set role and block
	* role 1 is user
	* block user (1 , 2) 2 is not block 1 is block
	 */
	user.Role = 1
	user.Block = 2
	/**
	* create new user based on register struct
	* token , role  , block will set with event
	 */
	err := a.repo.Create(user)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	// now user is login we can return his info
	helpers.OkResponse(g, "Thank you for register in our system", transformers.UserResponse(*user))
}

/**
* recover password take request token
* select user that have this token
* if user token valid and user not block
* then user can  recover his password
 */
func (a *AuthApp) Recover(g *gin.Context) {
	//init Reset struct to validate request
	recoverPassword := new(models.Recover)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if !auth.Recover(g, recoverPassword) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := a.repo.CheckUserExistsNotBlocked(g, "", recoverPassword.Token)
	if !valid {
		return
	}
	/**
	* now update token and update password
	* we update token to make it the old link not valid
	 */
	user.Password, _ = helpers.HashPassword(recoverPassword.Password)
	if user.Email == "admin@admin.com" {
		user.Token, _ = helpers.GenerateToken(user.Password + user.Email)
	}
	err := a.repo.Update(&user, user.UUID)
	if err != nil {
		helpers.ReturnForbidden(g, err.Error())
		return
	}
	// notice user that his password has been changes
	sendRecoverPasswordEmail(user, recoverPassword.Password)
	// return ok response
	helpers.OkResponse(g, "Your password has been set , and your token changes", transformers.UserResponse(user))
}

/***
* notice user that his password has been updated
 */
func sendRecoverPasswordEmail(user models.User, password string) {
	msg := "<h6>Your Password has been updated to (" + password + ")</h6>" + "<br>"
	msg += "<h6>Do not worry your password is encrypted , this just note for your activity</h6>"
	helpers.SendMail(user.Email, "Your password has been updated", msg)
}

/**
* reset password
* with email you can send reset link
* to user email
 */
func (a *AuthApp) Reset(g *gin.Context) {
	// init Reset struct to validate request
	reset := new(models.Reset)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if !auth.Reset(g, reset) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := a.repo.CheckUserExistsNotBlocked(g, reset.Email, "")
	if !valid {
		return
	}
	sendRestLink(user)
	// return ok response
	var data map[string]interface{}
	helpers.OkResponse(g, "We send your reset password link on your email", data)
}

/**
* create reset password link
* send it to user email
 */
func sendRestLink(user models.User) {
	msg := "<h6> Your Request To reset your password if you take this action click on this link to reset your password </h6>" + "<br>"
	msg += `<a href="` + os.Getenv("RESET_PASSWORD_URL") + user.Token + `">Reset Password</a>`
	helpers.SendMail(user.Email, "Reset Password Request", msg)
}
