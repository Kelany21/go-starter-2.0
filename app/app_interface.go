package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ApplicationInterface interface {
	Show(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func InitApps(db *gorm.DB) {
	NewUserApplication(db)
	NewAuthApplication(db)
}
