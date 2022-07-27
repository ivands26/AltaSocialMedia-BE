package delivery

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/common"
	"github.com/labstack/echo/v4"
)

type contentHandler struct {
	contentCases domain.ContentUseCases
}

func New(cs domain.ContentUseCases) domain.ContentHandler {
	return &contentHandler{
		contentCases: cs,
	}
}

func (cs *contentHandler) PostContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		//userId := common.ExtractData(c)
		var temp PostingFormat
		err := c.Bind(&temp)

		if err != nil {
			log.Println("Tidak bisa parse data", err)
			c.JSON(http.StatusBadRequest, "tidak bisa membaca input")
		}

		data, err := cs.contentCases.Posting(temp.ToModel())

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

func (cs *contentHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp PostingFormat
		result := c.Bind(&tmp)

		qry := map[string]interface{}{}
		id := common.ExtractData(c)

		if result != nil {
			log.Println("Cannot Parse Data", result)
			c.JSON(http.StatusBadRequest, "error read update")
		}

		if tmp.Content != "" {
			qry["content"] = tmp.Content
		}

		data, err := cs.contentCases.Update(id, tmp.ToModel())

		if err != nil {
			log.Println("Cannot Update Content", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"message": "Update Contect Success",
		})

	}
}

func (cs *contentHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := cs.contentCases.Delete(cnv)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete content",
		})
	}
}

func (cs *contentHandler) GetAllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := cs.contentCases.GetAllContent()

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("User Handler", err)
				c.JSON(http.StatusNotFound, err.Error())
			} else if strings.Contains(err.Error(), "retrieve") {
				log.Println("User Handler", err)
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all user data",
			"data":    data,
		})
	}
}
