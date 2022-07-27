package delivery

import (
	"log"
	"net/http"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/common"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCases
}

func New(us domain.UserUseCases) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
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

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		errLog := c.Bind(&userLogin)
		if errLog != nil {
			return c.JSON(http.StatusBadRequest, "email atau password salah")
		}
		username, token, err := us.userUsecase.Login(userLogin.Email, userLogin.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "login gagal")
		}
		data := map[string]interface{}{
			"username": username,
			"token":    token,
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login berhasil",
			"data":    data,
		})

	}

}
