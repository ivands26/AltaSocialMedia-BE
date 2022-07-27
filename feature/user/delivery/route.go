package delivery

import (
	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/user/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, bu domain.UserHandler) {
	e.POST("/login", bu.Login())
	e.POST("/register", bu.Register())
	e.GET("/profile", bu.GetSpecificUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
