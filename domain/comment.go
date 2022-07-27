package domain

import "github.com/labstack/echo/v4"

type Comment struct {
	ID        int
	Comment   string
	ContentID int
	UserID    int
}

// Error implements error
func (Comment) Error() string {
	panic("unimplemented")
}

type CommentHandler interface {
	GetAllComment() echo.HandlerFunc
	PostComment() echo.HandlerFunc
	// DeleteComment() echo.HandlerFunc
}

type CommentUseCases interface {
	PostingComment(userID int, newComment Comment) (Comment, error)
	GetAllComment() ([]Comment, error)
	// DeleteComment(commentID int) (Comment, error)
}

type DataComment interface {
	PostComment(newComment Comment) (Comment, error)
	GetAllComment() ([]Comment, error)
	// DeleteComment(commentID int) (Comment, error)
}
