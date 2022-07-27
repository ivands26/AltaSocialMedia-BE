package data

import (
	"github.com/AltaProject/AltaSocialMedia/domain"
	// "github.com/AltaProject/AltaSocialMedia/feature/comment/data"
	"github.com/AltaProject/AltaSocialMedia/feature/content/data"
	// "gorm.io/gorm"
)

type User struct {
	ID       int            `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
	Nama     string         `json:"nama" form:"nama" validate:"required"`
	Username string         `json:"username" form:"username" gorm:"unique"`
	Email    string         `json:"email" form:"email" validate:"required" gorm:"unique"`
	Password string         `json:"password" form:"password" validate:"required"`
	No_HP    string         `json:"no_hp" form:"no_hp"`
	Content  []data.Content `gorm:"foreignKey:UserID"`
	// Comment  []data.Comment `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       u.ID,
		Nama:     u.Nama,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		No_HP:    u.No_HP,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.User) User {
	var res User
	res.ID = data.ID
	res.Nama = data.Nama
	res.Username = data.Username
	res.Email = data.Email
	res.Password = data.Password
	res.No_HP = data.No_HP
	return res
}
