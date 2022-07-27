package data

import (
	"errors"
	"fmt"
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"gorm.io/gorm"
)

type ContentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ContentData {
	return &ContentData{
		db: db,
	}
}

func (cd *ContentData) AddNewContent(newContent domain.Content) (domain.Content, error) {
	var cnv = ToLocalContent(newContent)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("tidak bisa register", err.Error())
		return domain.Content{}, err
	}
	return cnv.toDomainContent(), nil
}

func (cd *ContentData) GetAllContent() ([]domain.Content, error) {
	var tmp []Content
	err := cd.db.Find(&tmp).Error

	if err != nil {
		log.Println("Cannot retrive object", err.Error())
		return nil, errors.New("cannot retrieve data")
	}

	if len(tmp) == 0 {
		log.Println("No data found", gorm.ErrRecordNotFound.Error())
		return nil, gorm.ErrRecordNotFound
	}

	return ParseArrDomainContent(tmp), nil
}

func (cd *ContentData) GetContentId(contenId int) (domain.Content, error) {
	var temp Content
	fmt.Println("contenId :", contenId)
	err := cd.db.Where("id = ?", contenId).First(&temp).Error
	if err != nil {
		log.Println("Data bermasalah / tidak ditemukan", err.Error())
		return domain.Content{}, err
	}
	fmt.Println("isi dari temp :", temp.toDomainContent())
	return temp.toDomainContent(), nil
}

func (cd *ContentData) Update(contentId int, newContent domain.Content) (domain.Content, error) {
	var content = ToLocalContent(newContent)
	err := cd.db.Model(&Content{}).Where("ID = ?", contentId).Updates(content)
	if err.Error != nil {
		log.Println("cant update content", err.Error.Error())
		return domain.Content{}, nil
	}

	if err.RowsAffected == 0 {
		log.Println("Content Not Updated")
		return domain.Content{}, nil

	}
	return content.toDomainContent(), nil

}

func (cd *ContentData) Delete(contentId int) bool {
	err := cd.db.Where("ID = ?", contentId).Delete(&Content{})
	if err.Error != nil {
		log.Println("cannot delete content", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No content deleted", err.Error.Error())
		return false
	}

	return true
}
