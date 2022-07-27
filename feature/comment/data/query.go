package data

import (
	"errors"
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"gorm.io/gorm"
)

type CommentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.DataComment {
	return &CommentData{
		db: db,
	}
}

func (cd *CommentData) GetAllComment() ([]domain.Comment, error) {
	var tmp []Comment
	err := cd.db.Find(&tmp).Error

	if err != nil {
		log.Println("Cannot retrive object", err.Error())
		return nil, errors.New("cannot retrieve data")
	}

	if len(tmp) == 0 {
		log.Println("No data found", gorm.ErrRecordNotFound.Error())
		return nil, gorm.ErrRecordNotFound
	}

	return ParseToArr(tmp), nil
}
