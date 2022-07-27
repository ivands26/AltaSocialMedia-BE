package data

import (
	"github.com/AltaProject/AltaSocialMedia/domain"
	// "github.com/AltaProject/AltaSocialMedia/feature/comment/data"
)

type Content struct {
	ID      int    `json:"id" form:"id" gorm:"primaryKey:autoIncrement"`
	Content string `json:"content" form:"content"`
	UserID  int
}

func (content *Content) toDomainContent() domain.Content {
	return domain.Content{
		ID:      int(content.ID),
		Content: content.Content,
		// UserID:  content.UserID,
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
	// res.UserID = data.UserID
	return res
}
