package delivery

import (
	"log"
	"net/http"
	"strings"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/labstack/echo/v4"
)

type commentHandler struct {
	commentCases domain.CommentUseCases
}

func New(cs domain.CommentUseCases) domain.CommentHandler {
	return &commentHandler{
		commentCases: cs,
	}
}

func (cs *commentHandler) GetAllComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := cs.commentCases.GetAllComment()
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
