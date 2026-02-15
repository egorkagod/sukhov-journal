package articles

import "context"

type ArticleCreateServiceDTO struct {
	AuthorID uint64
	Title    string
	Body     string
}

type ArticleEditServiceDTO struct {
	ID     uint64
	UserID uint64
	Title  string
	Body   string
}

type ArticleDeleteServiceDTO struct {
	ID     uint64
	UserID uint64
}

type ArticleService interface {
	GetRepo() ArticleRepo
	Create(ctx context.Context, data ArticleCreateServiceDTO) error
	Edit(ctx context.Context, data ArticleEditServiceDTO) error
	Delete(ctx context.Context, data ArticleDeleteServiceDTO) error
}

type articleService struct {
	repo ArticleRepo
}

func NewService(repo ArticleRepo) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) GetRepo() ArticleRepo {
	return s.repo
}

func (s *articleService) Create(ctx context.Context, data ArticleCreateServiceDTO) error {
	return s.repo.Create(ctx, ArticleCreateRepoDTO{AuthorID: data.AuthorID, Title: data.Title, Body: data.Body})
}

func (s *articleService) Edit(ctx context.Context, data ArticleEditServiceDTO) error {
	article, err := s.repo.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if article.AuthorID != data.UserID {
		return NoPermissionsErr
	}

	return s.repo.Edit(ctx, ArticleEditRepoDTO{ID: data.ID, Title: data.Title, Body: data.Body})
}

func (s *articleService) Delete(ctx context.Context, data ArticleDeleteServiceDTO) error {
	article, err := s.repo.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if article.AuthorID != data.UserID {
		return NoPermissionsErr
	}
	return s.repo.DeleteByID(ctx, data.ID)
}
