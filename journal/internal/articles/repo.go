package articles

import "context"


type ArticleCreateRepoDTO struct {
	AuthorID int64
	Title string
	Body string
}


type ArticleEditRepoDTO struct {
	ID int64
	Title string
	Body string
}


type ArticleRepo interface {
	GetByID(ctx context.Context, articleID int64) (Article, error)
	Create(ctx context.Context, data ArticleCreateRepoDTO) error
	Edit(ctx context.Context, data ArticleEditRepoDTO) error
	Delete(ctx context.Context, articleID int64) error
}