package factory

import (
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/utils"
	"gorm.io/gorm"
)

func seedUserAuth(db *gorm.DB) {
	adminUser := models.User{
		Username: "admin",
		Email:    "admin@localhost",
		Address:  "password",
		UserType: "salt",
	}

	isExist := db.Where("username = ?", adminUser.Username).First(&adminUser).RowsAffected
	if isExist == 0 {
		db.Create(&adminUser)
	}

	pass, salt := utils.Encrypt("password")
	adminAuth := models.Authentications{
		UserId:   adminUser.ID,
		Password: pass,
		Salt:     salt,
	}

	auth := db.Where("user_id = ?", adminUser.ID).First(&adminAuth).RowsAffected
	if auth == 0 {
		db.Create(&adminAuth)
	}
}
