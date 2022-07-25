package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       int
	Nama     string
	Username string
	Email    string
	Password string
	No_HP    string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	GetSpecificUser() echo.HandlerFunc
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
