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

func (cdd *CommentData) GetAllComment() ([]domain.Comment, error) {
	var tmp []Comment
	err := cdd.db.Find(&tmp).Error

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

func (cdd *CommentData) PostComment(newComment domain.Comment) (comment domain.Comment, err error) {
	var cnv = FromModel(newComment)
	err = cdd.db.Create(&cnv).Error
	if err != nil {
		log.Println("comment gagal diinput", err.Error())
		return domain.Comment{}, err
	}
	return cnv.ToModel(), nil
}
