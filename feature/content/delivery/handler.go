package delivery

import (
	"log"
	"net/http"

	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/common"
	"github.com/AltaProject/AltaSocialMedia/feature/content/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contentHandler struct {
	contentCases domain.ContentUseCases
}

func New(e *echo.Echo, cs domain.ContentUseCases) {
	handler := &contentHandler{
		contentCases: cs,
	}
	e.POST("/content/:id", handler.PostContent(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/content/:id", handler.GetSpecificContent(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}

func (cs *contentHandler) PostContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		//userId := common.ExtractData(c)
		var temp RegisterFormat
		err := c.Bind(&temp)

		if err != nil {
			log.Println("Tidak bisa parse data", err)
			c.JSON(http.StatusBadRequest, "tidak bisa membaca input")
		}

		// data, err := cs.contentCases.Posting(temp.ToModel())

		if err != nil {
			log.Println("tidak memproses data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "berhasil register data",
			"data":    data,
		})
	}
}

func (cs *contentHandler) GetSpecificContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := cs.contentCases.GetContentId(id)

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
