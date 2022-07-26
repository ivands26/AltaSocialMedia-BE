package delivery

import (
	"log"
	"net/http"

	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/common"
	"github.com/AltaProject/AltaSocialMedia/feature/user/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUseCases
}

func New(e *echo.Echo, us domain.UserUseCases) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.POST("/register", handler.Register())
	e.GET("/profile", handler.GetSpecificUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}

func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var temp RegisterFormat
		err := c.Bind(&temp)

		if err != nil {
			log.Println("Tidak bisa parse data", err)
			c.JSON(http.StatusBadRequest, "tidak bisa membaca input")
		}

		data, err := us.userUsecase.Register(temp.ToModel())

		if err != nil {
			log.Println("tidak memproses data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil register data",
			"data":    data,
			"token":   common.GenerateToken(data.ID),
		})
	}
}

func (us *userHandler) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := us.userUsecase.GetSpecificUser(id)

		if err != nil {
			log.Println("data tidak ditemukan", err)
			return c.JSON(http.StatusNotFound, "data tidak ditemukan")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "data ditemukan",
			"data":    data,
		})
	}
}
