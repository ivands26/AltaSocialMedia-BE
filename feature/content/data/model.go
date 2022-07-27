package data

import "github.com/AltaProject/AltaSocialMedia/domain"

type Content struct {
	ID      int    `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
	Content string `json:"content" form:"content" validate:"required"`
	UserID  int
	// User    User
	// Description string `json:"username" form:"username"`
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
