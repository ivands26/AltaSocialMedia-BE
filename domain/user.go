package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       int
	Nama     string `validate:"required"`
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	No_HP    string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	GetSpecificUser() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserUseCases interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
	Login(email string, password string) (username string, token string, err error)
}

type UserData interface {
	Register(newUser User) (User, error)
	GetSpecificUser(userId int) (User, error)
	Login(email string, password string) (username string, token string, err error)
}
