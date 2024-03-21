package middleware

import (
	"github.com/farhanaltariq/fiberplate/database"
	"github.com/farhanaltariq/fiberplate/database/redis"
	"github.com/farhanaltariq/fiberplate/services"
	"gorm.io/gorm"
)

type Services struct {
	DB                 *gorm.DB
	Rdb                *redis.Redis
	AuthService        services.AuthenticationService
	UserService        services.UserService
	ProductService     services.ProductService
	ProductTypeService services.ProductTypeService
}

func InitServices() Services {
	db := database.GetDBConnection()
	redis := redis.NewRedis("127.0.0.1:6379", "", 0)
	return Services{
		DB:                 db,
		Rdb:                redis,
		AuthService:        services.NewAuthService(db),
		UserService:        services.NewUserService(db),
		ProductService:     services.NewProductService(db),
		ProductTypeService: services.NewProductTypeService(db),
	}
}
