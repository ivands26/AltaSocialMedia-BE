package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type contentUseCase struct {
	dataContent domain.ContentData
	valid       *validator.Validate
}

func New(cd domain.ContentData, v *validator.Validate) domain.ContentUseCases {
	return &contentUseCase{
		dataContent: cd,
		valid:       v,
	}
}

func (cd *contentUseCase) Posting(userID int, newContent domain.Content) (domain.Content, error) {
	// var conv = data.FromModel(newUser)
	// err := ud.valid.Struct(conv)
	// if err != nil {
	// 	log.Println("Error Validasi", err)
	// 	return domain.User{}, err
	// }
	// hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Println("gagal enkripsi password", err)
	// 	return domain.User{}, err
	// }
	// newUser.Password = string(hashed)
	fmt.Println("id user :", userID)
	posting, err := cd.dataContent.AddNewContent(userID, newContent)

	if err != nil {
		log.Println(err.Error())
		return domain.Content{}, err
	}
	return posting, nil
}

func (cd *contentUseCase) GetContentId(contentId int) (domain.Content, error) {
	data, err := cd.dataContent.GetContentId(contentId)
	if err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.Content{}, errors.New("data tidak ditemukan")
		} else {
			return domain.Content{}, errors.New("server error")
		}
	}
	return data, nil
}

func (cd *contentUseCase) Update(contentId int, newContent domain.Content) (domain.Content, error) {
	UpdContent := map[string]interface{}{}
	if newContent.Content != "" {
		UpdContent["content"] = newContent.Content
	}

	update, err := cd.dataContent.Update(contentId, newContent)

	if err != nil {
		log.Println(err.Error())
		return domain.Content{}, err
	}

	return update, nil
}

func (cd *contentUseCase) Delete(contentId int) (bool, error) {
	res := cd.dataContent.Delete(contentId)

	if !res {
		return false, errors.New("failed to delete content")
	}
	return true, nil
}
