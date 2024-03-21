package middleware

import (
	"github.com/farhanaltariq/fiberplate/database"
	"github.com/farhanaltariq/fiberplate/services"
	"gorm.io/gorm"
)

type Services struct {
	DB                 *gorm.DB
	AuthService        services.AuthenticationService
	UserService        services.UserService
	ProductService     services.ProductService
	ProductTypeService services.ProductTypeService
}

func InitServices() Services {
	db := database.GetDBConnection()
	return Services{
		DB:                 db,
		AuthService:        services.NewAuthService(db),
		UserService:        services.NewUserService(db),
		ProductService:     services.NewProductService(db),
		ProductTypeService: services.NewProductTypeService(db),
	}
}
