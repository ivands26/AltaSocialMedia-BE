package factory

import (
	ud "github.com/AltaProject/AltaSocialMedia/feature/user/data"
	userDelivery "github.com/AltaProject/AltaSocialMedia/feature/user/delivery"
	us "github.com/AltaProject/AltaSocialMedia/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	valid := validator.New()
	useCase := us.New(userData, valid)
	userDelivery.New(e, useCase)

}
