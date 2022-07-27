package delivery

import "github.com/AltaProject/AltaSocialMedia/domain"

type CommentFormat struct {
	Comment string `json:"comment"`
}

func (p *CommentFormat) ToModel() domain.Comment {
	return domain.Comment{
		Comment: p.Comment,
	}
}
