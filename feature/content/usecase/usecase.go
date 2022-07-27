package usecase

import (
	"errors"
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
	if userID == -1 {
		return domain.Content{}, errors.New("invalid user")
	}

	newContent.UserID = userID
	posting, err := cd.dataContent.AddNewContent(newContent)

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

func (cd *contentUseCase) GetAllContent() ([]domain.Content, error) {
	data, err := cd.dataContent.GetAllContent()

	if err == gorm.ErrRecordNotFound {
		log.Println("User Usecase", err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Println("User Usecase", err.Error())
		return nil, errors.New("error when retrieve data")
	}
	return data, nil
}
