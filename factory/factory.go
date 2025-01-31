package factory

import (
	ud "github.com/AltaProject/AltaSocialMedia/feature/user/data"
	uDelivery "github.com/AltaProject/AltaSocialMedia/feature/user/delivery"
	us "github.com/AltaProject/AltaSocialMedia/feature/user/usecase"

	cd "github.com/AltaProject/AltaSocialMedia/feature/content/data"
	cDelivery "github.com/AltaProject/AltaSocialMedia/feature/content/delivery"
	cs "github.com/AltaProject/AltaSocialMedia/feature/content/usecase"

	commData "github.com/AltaProject/AltaSocialMedia/feature/content/data"
	commDeliver "github.com/AltaProject/AltaSocialMedia/feature/content/delivery"
	commCase "github.com/AltaProject/AltaSocialMedia/feature/content/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	valid := validator.New()
	userCase := us.New(userData, valid)
	userHandler := uDelivery.New(userCase)
	uDelivery.RouteUser(e, userHandler)

	cData := cd.New(db)
	cCase := cs.New(cData, valid)
	cHandler := cDelivery.New(cCase)
	cDelivery.RouteContent(e, cHandler)

	comD := commData.New(db)
	comC := commCase.New(comD, valid)
	comHandler := commDeliver.New(comC)
	cDelivery.RouteContent(e, comHandler)

}
