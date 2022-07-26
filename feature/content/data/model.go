package data

import "github.com/AltaProject/AltaSocialMedia/domain"

type Content struct {
	ID      int    `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
	Content string `json:"nama" form:"nama" validate:"required"`
	// Description string `json:"username" form:"username"`
	UserID int
}

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
	var res Content
	res.ID = int(data.ID)
	res.Content = data.Content
	res.UserID = data.UserID
	return res
}
