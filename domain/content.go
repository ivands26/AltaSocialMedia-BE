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
	GetSpecificContent() echo.HandlerFunc
}

type ContentUseCases interface {
	Posting(newContent Content) (Content, error)
	GetContentId(contentId int) (Content, error)
}

type ContentData interface {
	AddNewContent(newContent Content) (Content, error)
	GetContentId(contentId int) (Content, error)
}
