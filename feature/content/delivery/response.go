package delivery

import "github.com/AltaProject/AltaSocialMedia/domain"

type ContentRes struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func FromDomain(data domain.Content) ContentRes {
	return ContentRes{
		ID:      int(data.ID),
		Content: data.Content,
		UserID:  int(data.UserID),
	}
}
