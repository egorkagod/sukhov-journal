package articles

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ArticleCreateRepoDTO struct {
	AuthorID uint64
	Title    string
	Body     string
}

type ArticleEditRepoDTO struct {
	ID    uint64
	Title string
	Body  string
}

type ArticleRepo interface {
	GetByID(ctx context.Context, id uint64) (*Article, error)
	Create(ctx context.Context, data ArticleCreateRepoDTO) error
	Edit(ctx context.Context, data ArticleEditRepoDTO) error
}

type articleRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) GetByID(ctx context.Context, id uint64) (*Article, error) {
	var article Article
	query := r.db.WithContext(ctx).First(&article, id)
	err := query.Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ArticleNotFoundErr
	case err != nil:
		return nil, err
	}

	return &article, nil
}

func (r *articleRepo) Create(ctx context.Context, data ArticleCreateRepoDTO) error {
	newArticle := &Article{AuthorID: data.AuthorID, Title: data.Title, Body: data.Body}
	return r.db.WithContext(ctx).Create(newArticle).Error
}

func (r *articleRepo) Edit(ctx context.Context, data ArticleEditRepoDTO) error {
	article, err := r.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	article.Title = data.Title
	article.Body = data.Body
	tx := r.db.Save(article)
	return tx.Error
}
