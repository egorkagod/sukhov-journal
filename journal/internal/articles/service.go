package articles

import "context"


type ArticleCreateServiceDTO struct {
	AuthorID int64
	Title string
	Body string
}


type ArticleEditServiceDTO struct {
	ID int64
	UserID int64
	Title string
	Body string
}

type ArticleService interface {
	Create(ctx context.Context, data ArticleCreateServiceDTO) error
	Edit(ctx context.Context, data ArticleEditServiceDTO) error
	Delete(ctx context.Context, articleID int64) error
}