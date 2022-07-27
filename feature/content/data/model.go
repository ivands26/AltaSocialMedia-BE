package data

import (
	"github.com/AltaProject/AltaSocialMedia/domain"
	// "github.com/AltaProject/AltaSocialMedia/feature/comment/data"
)

type Content struct {
	ID      int    `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
<<<<<<< HEAD
	Content string `json:"content" form:"content" validate:"required"`
	UserID  int
	// User    User
	// Description string `json:"username" form:"username"`
=======
	Content string `json:"content" form:"content"`
	UserID  int
>>>>>>> ac04ca53f27afc5f71e6338e4b1c9b9f0322840d
}

// type User struct {
// 	ID       uint
// 	Nama     string
// 	Username string
// 	Content  []Content `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
// }

func (content *Content) toDomainContent() domain.Content {
	return domain.Content{
		ID:      int(content.ID),
		Content: content.Content,
		UserID:  content.UserID,
	}
}

func ParseArrDomainContent(arr []Content) []domain.Content {
	var res []domain.Content

	for _, val := range arr {
		res = append(res, val.toDomainContent())
	}
	return res
}

func ToLocalContent(data domain.Content) Content {
	return Content{
		ID:      int(data.ID),
		Content: data.Content,
		UserID:  data.UserID,
		// User: User{
		// 	Nama: data.User.Name,
		// 	Username: User.Username,

		// },
	}
}

// var res Content
// res.ID = int(data.ID)
// res.Content = data.Content
// res.UserID = data.UserID
// res.User.Nama =
// return res
// }
