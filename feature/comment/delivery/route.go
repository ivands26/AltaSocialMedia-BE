package delivery

import (
	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/content/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteComment(e *echo.Echo, ch domain.CommentHandler) {
	e.GET("/comment/:id", ch.GetAllComment(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))

}
