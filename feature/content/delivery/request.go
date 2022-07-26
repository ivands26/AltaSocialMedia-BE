package delivery

import "github.com/AltaProject/AltaSocialMedia/domain"

type PostingFormat struct {
	Content string `json:"no_hp"`
}

func (p *PostingFormat) ToModel() domain.Content {
	return domain.Content{
		Content: p.Content,
	}
}
