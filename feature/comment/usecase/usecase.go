package usecase

import (
	"errors"
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type commentUseCase struct {
	dataComment domain.DataComment
	valid       *validator.Validate
}

func New(cd domain.DataComment, v *validator.Validate) domain.CommentUseCases {
	return &commentUseCase{
		dataComment: cd,
		valid:       v,
	}
}

func (cd *commentUseCase) GetAllComment() ([]domain.Comment, error) {
	data, err := cd.dataComment.GetAllComment()

	if err == gorm.ErrRecordNotFound {
		log.Println("User Usecase", err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Println("User Usecase", err.Error())
		return nil, errors.New("error when retrieve data")
	}
	return data, nil
}

func (cd *commentUseCase) PostingComment(userID int, newComment domain.Comment) (domain.Comment, error) {
	if userID == -1 {
		return domain.Comment{}, errors.New("invalid user")
	}

	newComment.UserID = userID
	posting, err := cd.dataComment.PostComment(newComment)

	if err != nil {
		log.Println(err.Error())
		return domain.Comment{}, err
	}
	return posting, nil
}
