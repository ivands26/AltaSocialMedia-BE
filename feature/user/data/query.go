package data

import (
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Register(newUser domain.User) (domain.User, error) {
	var cnv = FromModel(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("tidak bisa register", err.Error())
		return domain.User{}, err
	}
	return cnv.ToModel(), nil
}

func (ud *userData) GetSpecificUser(userId int) (domain.User, error) {
	var temp User
	err := ud.db.Where("ID = ?", userId).First(&temp).Error
	if err != nil {
		log.Println("Data bermasalah / tidak ditemukan", err.Error())
		return domain.User{}, err
	}
	return temp.ToModel(), nil
}
