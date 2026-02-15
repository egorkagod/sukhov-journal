package articles

import (
	"context"
	"errors"
	"time"

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
	DeleteByID(ctx context.Context, id uint64) error
}

type articleRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) GetByID(ctx context.Context, id uint64) (*Article, error) {
	var article Article
	query := r.db.WithContext(ctx).Where("id = ? AND is_deleted = false", id).First(&article)
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
	newArticle := &Article{AuthorID: data.AuthorID, Title: data.Title, Body: data.Body, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	return r.db.WithContext(ctx).Create(newArticle).Error
}

func (r *articleRepo) Edit(ctx context.Context, data ArticleEditRepoDTO) error {
	article, err := r.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	article.Title = data.Title
	article.Body = data.Body
	article.UpdatedAt = time.Now()
	tx := r.db.Save(article)
	return tx.Error
}

func (r *articleRepo) DeleteByID(ctx context.Context, id uint64) error {
	article, err := r.GetByID(ctx, id)
	if err != nil {
		return err
	}

	article.IsDeleted = true
	tx := r.db.Save(article)
	return tx.Error
}
