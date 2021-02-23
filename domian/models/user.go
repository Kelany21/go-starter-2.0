package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang-ddd-starter/helpers"
)

/***
* model struct here we will build the main
* struct that connect to database
 */
type User struct {
	Model
	Name     string `gorm:"type:varchar(50);" json:"name"`
	Email    string `gorm:"type:varchar(50);unique_index" json:"email"`
	Role     int    `gorm:"_" json:"role"`
	Password string `gorm:"size:255" json:"-"`
	Token    string `gorm:"size:255" json:"token"`
	Block    int    `gorm:"_" json:"block"`
}

//UserApp implements the UserApplication interface
var _ helpers.JsonApiInterface = &User{}

func (user User) GetID() string {
	return user.UUID.String()
}

func (user User) GetType() string {
	return "user"
}

/**
* user included data
 */
func (user User) GetIncludes() map[string]interface{} {
	var u = make(map[string]interface{})
	u["action"] = UserActions(user)

	return u
}

/**
* user included actions
 */
func UserActions(user User) []map[string]interface{} {
	u := []map[string]interface{}{
		{
			"id": uuid.New(),
			"attributes": map[string]interface{}{
				"endpoint_url": ";jobnvdf" + user.UUID.String(),
				"method":       "GET",
				"label":        "vd",
				"key":          "vcxv",
			},
			"relationships": nil,
		},
	}

	return u
}

/**
* use this struct when visitor login
 */
type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

/**
* use this struct when visitor login
 */
type Logout struct {
	Token string `json:"token"`
}

/**
* use this struct when reset email
 */
type Reset struct {
	Email string `json:"email"`
}

/**
* use this struct when reset email
 */
type Recover struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

/**
* event when user register
* create token
* hash password
 */
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	if user.Email != "admin@admin.com" {
		user.UUID = uuid.New()
		token, _ := helpers.HashPassword(user.Email + user.Password)
		password, _ := helpers.HashPassword(user.Password)
		scope.SetColumn("token", token)
		scope.SetColumn("password", password)
	}

	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) (err error) {
	password, _ := helpers.HashPassword(user.Password)
	scope.SetColumn("password", password)
	if user.Email != "admin@admin.com" {
		token, _ := helpers.HashPassword(user.Email + user.Password)
		scope.SetColumn("token", token)
	}
	return
}

/**
* you can update these column only
 */
func UserFillAbleColumn() []string {
	return []string{"name", "email", "role", "password", "block"}
}

/**
* active category only
 */
func ActiveUser(db *gorm.DB) *gorm.DB {
	return db.Where("status = 2")
}
