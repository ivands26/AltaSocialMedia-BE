package data

import (
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"gorm.io/gorm"
)

type contentData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ContentData {
	return &contentData{
		db: db,
	}
}

func (cd *contentData) AddNewContent(newContent domain.Content) (domain.Content, error) {
	var cnv = ToLocalContent(newContent)
	err := cd.db.Create(&cnv).Error
	if err != nil {
		log.Println("tidak bisa register", err.Error())
		return domain.Content{}, err
	}
	return cnv.toDomainContent(), nil
}

func (cd *contentData) GetContentId(contenId int) (domain.Content, error) {
	var temp Content
	err := cd.db.Where("ID = ?", contenId).First(&temp).Error
	if err != nil {
		log.Println("Data bermasalah / tidak ditemukan", err.Error())
		return domain.Content{}, err
	}
	return temp.toDomainContent(), nil
}

func (cd *contentData) Update(userId int, newContent domain.Content) (domain.Content, error) {
	var content = ToLocalContent(newContent)
	err := cd.db.Model(&Content{}).Where("ID = ?", content.ID).Updates(content)
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
