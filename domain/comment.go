package domain

import "github.com/labstack/echo/v4"

type Comment struct {
	ID        int
	Comment   string
	ContentID int
	UserID    int
}

type CommentHandler interface {
	GetAllComment() echo.HandlerFunc
	// PostCommentId() echo.HandlerFunc
	// DeleteComment() echo.HandlerFunc
}

type CommentUseCases interface {
	// InsertComment(newComment Comment, contentID int) (Comment, error)
	GetAllComment() ([]Comment, error)
	// DeleteComment(commentID int) (Comment, error)
}

type DataComment interface {
	// AddComment(newComment Comment, contentID int) (Comment, error)
	GetAllComment() ([]Comment, error)
	// DeleteComment(commentID int) (Comment, error)
}
