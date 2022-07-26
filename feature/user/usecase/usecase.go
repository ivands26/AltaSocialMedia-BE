package usecase

import (
	"errors"
	"log"

	"github.com/AltaProject/AltaSocialMedia/domain"
	// "github.com/AltaProject/AltaSocialMedia/feature/user/data"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userCase struct {
	userData domain.UserData
	valid    *validator.Validate
}

func New(ud domain.UserData, val *validator.Validate) domain.UserUseCases {
	return &userCase{
		userData: ud,
		valid:    val,
	}
}

func (ud *userCase) Register(newUser domain.User) (domain.User, error) {
	// var conv = data.FromModel(newUser)
	// err := ud.valid.Struct(conv)
	// if err != nil {
	// 	log.Println("Error Validasi", err)
	// 	return domain.User{}, err
	// }
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("gagal enkripsi password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	register, err := ud.userData.Register(newUser)

	if err != nil {
		log.Println(err.Error())
		return domain.User{}, err
	}

	if register.ID == 0 {
		return domain.User{}, errors.New("tidak registrasi")
	}

	return register, nil
}

func (ud *userCase) GetSpecificUser(userId int) (domain.User, error) {
	data, err := ud.userData.GetSpecificUser(userId)
	if err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data tidak ditemukan")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (ud *userCase) Login(email string, password string) (username string, token string, err error) {
	username, token, err = ud.userData.Login(email, password)
	return username, token, err
}
