package database

import (
	model "github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Database struct {
	*gorm.DB
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Authentications{},
	)
}

func Connect() error {
	dsn := utils.GetEnv("DB_DSN", "user=postgres password=postgres dbname=boilerplate port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		logrus.Errorln("Error connecting to database", err)
		return err
	}

	AutoMigrate(db)
	SetupDBConnection(db)
	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
