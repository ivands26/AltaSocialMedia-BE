package delivery

import (
	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/content/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteContent(e *echo.Echo, ch domain.ContentHandler) {
	e.POST("/content", ch.PostContent(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/content/:id", ch.Update(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/content/:id", ch.Delete(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/content/:id", ch.GetSpecificContent(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/content", ch.GetSpecificContent())

}
