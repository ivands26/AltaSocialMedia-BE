package delivery

// import (
// 	"github.com/AltaProject/AltaSocialMedia/domain"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func RoutesUser(e *echo.Echo, uc domain.UserHandler) {
// 	// e.Pre(middleware.RemoveTrailingSlash())

// 	// e.Use(middleware.CORS())

// 	e.POST("/user", uc.Register())
// 	e.GET("/user/:userID", uc.GetSpecificUser(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("ALTASOSMED")}))
// }
