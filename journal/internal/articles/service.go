package articles

import (
	"context"
	"log"

	"journal/internal/voice"
)

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
	GetAudioPath(ctx context.Context, articleID uint64) (string, error)
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
	articleID, err := s.repo.Create(ctx, ArticleCreateRepoDTO{AuthorID: data.AuthorID, Title: data.Title, Body: data.Body})
	if err != nil {
		return err
	}
	go s.VoiceOver(context.Background(), articleID)
	return nil
}

func (s *articleService) Edit(ctx context.Context, data ArticleEditServiceDTO) error {
	article, err := s.repo.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if article.AuthorID != data.UserID {
		return ErrNoPermision
	}

	return s.repo.Edit(ctx, ArticleEditRepoDTO{ID: data.ID, Title: data.Title, Body: data.Body})
}

func (s *articleService) Delete(ctx context.Context, data ArticleDeleteServiceDTO) error {
	article, err := s.repo.GetByID(ctx, data.ID)
	if err != nil {
		return err
	}

	if article.AuthorID != data.UserID {
		return ErrNoPermision
	}
	return s.repo.DeleteByID(ctx, data.ID)
}

func (s *articleService) GetAudioPath(ctx context.Context, articleID uint64) (string, error) {
	article, err := s.repo.GetByID(ctx, articleID)
	if err != nil {
		return "", err
	}

	return article.AudioPath, nil
}

func (s *articleService) VoiceOver(ctx context.Context, articleID uint64) {
	article, err := s.repo.GetByID(ctx, articleID)
	if err != nil {
		log.Printf("Ошибка при получении статьи для озвучки - %v", err.Error())
		return
	}

	audioData, err := voice.Manager.VoiceOver(article.Body)
	if err != nil {
		log.Printf("Ошибка при получении аудио озвучки статьи - %v", err.Error())
		return
	}

	path, err := voice.Manager.SaveAudio(audioData, articleID)
	if err != nil {
		log.Printf("Ошибка при сохранении аудио озвучки статьи - %v", err.Error())
		return
	}

	err = s.repo.AddAudioPath(ctx, ArticleAddAudioPathDTO{ID: articleID, AudioPath: path})
	if err != nil {
		log.Printf("Ошибка при сохранении пути к аудио в статье - %v", err.Error())
		return
	}
}
