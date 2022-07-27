package domain

import "github.com/labstack/echo/v4"

type Content struct {
	ID      int
	Content string
	UserID  int
	// Description string
}

type ContentHandler interface {
	PostContent() echo.HandlerFunc
	GetAllContent() echo.HandlerFunc
	GetSpecificContent() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ContentUseCases interface {
	Posting(userID int, newContent Content) (Content, error)
	GetContentId(contentId int) (Content, error)
	GetAllContent() ([]Content, error)
	Update(userId int, newContent Content) (Content, error)
	Delete(contentId int) (bool, error)
}

type ContentData interface {
	AddNewContent(newContent Content) (Content, error)
	GetAllContent() ([]Content, error)
	GetContentId(contentId int) (Content, error)
	Update(userId int, newContent Content) (Content, error)
	Delete(contentId int) bool
}
