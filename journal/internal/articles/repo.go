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

type ArticleAddAudioPathDTO struct {
	ID        uint64
	AudioPath string
}

type ArticleRepo interface {
	GetByID(ctx context.Context, id uint64) (*Article, error)
	Create(ctx context.Context, data ArticleCreateRepoDTO) (uint64, error)
	Edit(ctx context.Context, data ArticleEditRepoDTO) error
	DeleteByID(ctx context.Context, id uint64) error
	AddAudioPath(ctx context.Context, data ArticleAddAudioPathDTO) error
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
		return nil, ErrArticleNotFound
	case err != nil:
		return nil, err
	}

	return &article, nil
}

func (r *articleRepo) Create(ctx context.Context, data ArticleCreateRepoDTO) (uint64, error) {
	newArticle := &Article{AuthorID: data.AuthorID, Title: data.Title, Body: data.Body, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := r.db.WithContext(ctx).Create(newArticle).Error
	if err != nil {
		return 0, err
	}
	return newArticle.ID, nil
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

func (r *articleRepo) AddAudioPath(ctx context.Context, data ArticleAddAudioPathDTO) error {
	article, err := r.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	article.AudioPath = data.AudioPath
	tx := r.db.Save(article)
	return tx.Error
}
