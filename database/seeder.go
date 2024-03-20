package database

import (
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/utils"
)

func seedUserAuth() {
	adminUser := models.User{
		Username: "admin",
		Email:    "admin@localhost",
		Address:  "password",
		UserType: "salt",
	}

	isExist := db.Debug().Where("username = ?", adminUser.Username).First(&adminUser).RowsAffected
	if isExist == 0 {
		db.Debug().Create(&adminUser)
	}

	pass, salt := utils.Encrypt("password")
	adminAuth := models.Authentications{
		UserId:   adminUser.ID,
		Password: pass,
		Salt:     salt,
	}

	auth := db.Debug().Where("user_id = ?", adminUser.ID).First(&adminAuth).RowsAffected
	if auth == 0 {
		db.Debug().Create(&adminAuth)
	}
}

func Seed() {
	seedUserAuth()
}
