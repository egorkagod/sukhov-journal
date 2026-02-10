package articles

import "context"

type ArticleCreateRepoDTO struct {
	AuthorID int64
	Title    string
	Body     string
}

type ArticleEditRepoDTO struct {
	ID    int64
	Title string
	Body  string
}

type ArticleRepo interface {
	GetByID(ctx context.Context, id int64) (*Article, error)
	Create(ctx context.Context, data ArticleCreateRepoDTO) error
	Edit(ctx context.Context, data ArticleEditRepoDTO) error
}

type articleInMemoryRepo struct {
	data []Article
}

func NewRepo() ArticleRepo {
	return &articleInMemoryRepo{data: make([]Article, 0)}
}

func (r *articleInMemoryRepo) GetByID(ctx context.Context, id int64) (*Article, error) {
	if id >= 0 && id < int64(len(r.data)) {
		return &r.data[id], nil
	}
	return nil, ArticleNotFoundErr
}

func (r *articleInMemoryRepo) Create(ctx context.Context, data ArticleCreateRepoDTO) error {
	newArticle := Article{ID: int64(len(r.data)), AuthorID: data.AuthorID, Title: data.Title, Body: data.Body}
	r.data = append(r.data, newArticle)
	return nil
}

func (r *articleInMemoryRepo) Edit(ctx context.Context, data ArticleEditRepoDTO) error {
	previousArticle, err := r.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	newArticle := Article{ID: data.ID, AuthorID: previousArticle.AuthorID, Title: data.Title, Body: data.Body}
	r.data[data.ID] = newArticle
	return nil
}
